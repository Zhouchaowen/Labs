//go:build linux
// +build linux

package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"runtime"
)

type eventList struct {
	size   int
	events []unix.EpollEvent
}

func newEventList(size int) *eventList {
	return &eventList{size, make([]unix.EpollEvent, size)}
}

func main() {
	run()
}

func run() {
	// 获取是tcp的listenFd
	ListenFd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM|unix.SOCK_CLOEXEC, unix.IPPROTO_TCP)
	if err != nil {
		fmt.Println("create socket err", err)
		return
	}
	err = unix.SetsockoptInt(ListenFd, unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
	if err != nil {
		fmt.Println("set socket err", err)
		return
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println("resove err", err)
		return
	}
	sa := &unix.SockaddrInet4{
		Port: tcpAddr.Port,
	}
	// 绑定的端口
	err = unix.Bind(ListenFd, sa)
	if err != nil {
		fmt.Println("socket bind err", err)
		return
	}
	var n int
	if n > 1<<16-1 {
		n = 1<<16 - 1
	}
	// 监听服务
	err = unix.Listen(ListenFd, n)
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	// 开启epoll
	el := newEventList(128)
	epollFd, err := unix.EpollCreate1(unix.EPOLL_CLOEXEC)
	if err != nil {
		unix.Close(epollFd)
		fmt.Println("create fd err", err)
		return
	}

	go func() {
		buf := make([]byte, 1024)
		var msec = -1
		for {
			// 有读写事件到来就会得到通知
			nready, err := unix.EpollWait(epollFd, el.events, msec)
			if nready <= 0 {
				msec = -1
				runtime.Gosched()
				continue
			}
			msec = 0
			if err != nil {
				if err == unix.EINTR {
					continue
				}
				fmt.Println("listenner wait err", err)
				return
			}
			for i := 0; i < nready; i++ {
				if el.events[i].Events&unix.EPOLLERR|unix.EPOLLHUP|unix.EPOLLRDHUP|unix.EPOLLIN|unix.EPOLLPRI != 0 {
					fd := el.events[i].Fd
					n, err := unix.Read(int(fd), buf)
					if err != nil {
						unix.Close(int(el.events[i].Fd))
						break
					}
					if n > 0 {
						ev := unix.EpollEvent{
							Events: unix.EPOLLOUT,
							Fd:     fd,
						}
						if err := unix.EpollCtl(epollFd, unix.EPOLL_CTL_MOD, int(fd), &ev); err != nil {
							fmt.Println("get data   err")
							continue
						}
						fmt.Println("[suc]", string(buf[:n]))
					}
					if n <= 0 {
						ev := unix.EpollEvent{
							Events: unix.EPOLLOUT | unix.EPOLLIN | unix.EPOLLERR | unix.EPOLLHUP,
							Fd:     fd,
						}
						if err := unix.EpollCtl(epollFd, unix.EPOLL_CTL_DEL, int(fd), &ev); err != nil {
							fmt.Println("close epoll err")
							return
						}
						unix.Close(int(fd))
					}
				}
			}
		}
	}()
	for {
		// 获取连接
		conn, _, err := unix.Accept(ListenFd)
		if err != nil {
			fmt.Println("accept err", err)
			return
		}
		err = unix.SetNonblock(conn, true)
		if err != nil {
			fmt.Println("block err", err)
			return
		}
		ev := unix.EpollEvent{}
		ev.Fd = int32(conn)
		ev.Events = unix.EPOLLPRI | unix.EPOLLIN
		// 把链接注册到epoll中
		err = unix.EpollCtl(epollFd, unix.EPOLL_CTL_ADD, conn, &ev)
		if err != nil {
			fmt.Println("epoll ctl err", err)
			return
		}
	}
}
