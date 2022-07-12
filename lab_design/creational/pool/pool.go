/*
1.当创建类对象的成本很高并且在特定时间需要的此类对象的数量不多时
	让我们以 DB 连接为例。每个连接对象的创建成本都很高，因为涉及到网络调用，而且在某个时候可能不需要超过某个连接。对象池设计模式非常适合这种情况。
2.当池对象是不可变对象时
	再次以数据库连接为例。数据库连接是一个不可变对象。几乎不需要更改它的所有属性
3.出于性能原因。由于池已经创建，它将显着提高应用程序性能
*/
// https://golangbyexample.com/golang-object-pool/
package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

type iPoolObject interface {
	getID() string //This is any id which can be used to compare two different pool objects
}

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mulock   *sync.Mutex
}

//InitPool Initialize the pool
func initPool(poolObjects []iPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Cannot craete a pool of 0 length")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}

func (p *pool) loan() (iPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free. Please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())
	return obj, nil
}

func (p *pool) receive(target iPoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
	return nil
}

func (p *pool) remove(target iPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("Targe pool object doesn't belong to the pool")
}

type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}

func main() {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := initPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	conn1, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn2, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	pool.receive(conn1)
	pool.receive(conn2)
}
