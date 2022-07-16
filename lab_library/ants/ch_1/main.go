//注意
//线程池执行有两种，一种执行普通逻辑方法pool，可接受所有方法,另一种执行形同类型的方法（就是每次接收的内容方法都一样）
//使用前需要先建立一个对应的pool对象，参数是容量大小和过期时间等， 如果使用普通方法，有默认方法可以使用
//使用后，需要调用Release来结束使用，意思就是close那类的意思
//总的来说，使用过程就是，新建对象，加入逻辑函数代码，结束使用关闭对象。
package main

import (
	"log"
	"sync"
	"time"

	"github.com/panjf2000/ants"
)

const (
	runTime = 10000
)

//使用默认普通pool
//其实就是使用了普通的pool，为了方便直接使用，在内部已经new了一个普通的pool，
//相当于下面那个新建的过程给你写好了，容量大小和过期时间都用默认的，详细信息可以看源码，里面剥一层就可以看到
func main() {
	defer ants.Release() //退出工作，相当于使用后关闭

	var wg sync.WaitGroup //这里使用等待是为了看出结果，阻塞主线程，防止直接停止，如果在web项目中，就不需要
	log.Println("start ants work")
	for i := 0; i < runTime; i++ {
		wg.Add(1)
		ants.Submit(func() { //提交函数，将逻辑函数提交至work中执行，这里写入自己的逻辑
			//log.Println(i, ":hello")
			time.Sleep(time.Millisecond * 10)
			wg.Done()
		})
	}
	wg.Wait()
	log.Printf("pool, capacity:%d", ants.Cap())
	log.Printf("pool, running workers number:%d", ants.Running())
	log.Printf("pool, free workers number:%d", ants.Free())
}
