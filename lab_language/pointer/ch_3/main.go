// Pointer和Value方法的区别(Pointer地址不会变，Value随用随分配)
package main

import "fmt"

/*
	Pointer 使用指针方法当前地址固定
	Value   使用值方法当前地址随用随分配

	addr P:0xc00000c030
	addr P.A:0xc00000c030
	addr P.B:0xc00000c040
	addr P:0xc00000c030
	addr P.A:0xc00000c030
	addr P.B:0xc00000c040
	-----------
	addr V:0xc00000c048
	addr V.A:0xc00000c048
	addr V.B:0xc00000c058
	addr V:0xc00000c060
	addr V.A:0xc00000c060
	addr V.B:0xc00000c070
*/

type Pointer struct {
	A string
	B int
}

func (p *Pointer) methodA() {
	fmt.Printf("addr P:%p\n", p)
	fmt.Printf("addr P.A:%p\n", &p.A)
	fmt.Printf("addr P.B:%p\n", &p.B)
}

func (p *Pointer) methodB() {
	fmt.Printf("addr P:%p\n", p)
	fmt.Printf("addr P.A:%p\n", &p.A)
	fmt.Printf("addr P.B:%p\n", &p.B)
}

type Value struct {
	A string
	B int
}

func (v Value) methodA() {
	fmt.Printf("addr V:%p\n", &v)
	fmt.Printf("addr V.A:%p\n", &v.A)
	fmt.Printf("addr V.B:%p\n", &v.B)
}

func (v Value) methodB() {
	fmt.Printf("addr V:%p\n", &v)
	fmt.Printf("addr V.A:%p\n", &v.A)
	fmt.Printf("addr V.B:%p\n", &v.B)
}

func main() {
	p := Pointer{}
	p.methodA()
	p.methodB()

	fmt.Println("-----------")
	v := Value{}
	v.methodA()
	v.methodB()
}
