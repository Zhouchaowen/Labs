package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	Age  int
}

func func1() {
	var a int = 50
	v := reflect.ValueOf(a) // 返回Value类型对象，值为50
	t := reflect.TypeOf(a)  // 返回Type类型对象，值为int
	// 50 int int int 50
	fmt.Println(v, t, v.Type(), t.Kind(), reflect.ValueOf(&a).Elem())
	seta := reflect.ValueOf(&a).Elem() // 这样才能让seta保存a的值
	// 50 true
	fmt.Println(seta, seta.CanSet())
	seta.SetInt(1000)
	// seta:  1000
	// a:  1000
	fmt.Println("seta: ", seta)
	fmt.Println("a: ", a)
}

func func2() {
	var b [5]int = [5]int{5, 6, 7, 8}
	// [5]int array int
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem())
	bb := reflect.ValueOf(&b).Elem()
	// [5 6 7 8 0] true
	fmt.Println(bb, bb.CanSet())
	bb.Index(0).SetInt(10)
	// [10 6 7 8 0]
	fmt.Println(bb)
}

func func3() {
	var Pupil Student = Student{"joke", 18}
	p := reflect.ValueOf(Pupil) // 使用ValueOf()获取到结构体的Value对象
	// main.Student
	fmt.Println(p.Type()) // 输出:Student
	// struct
	fmt.Println(p.Kind()) // 输出:struct

	setStudent := reflect.ValueOf(&Pupil).Elem()
	//setStudent.Field(0).SetString("Mike") // 未导出字段，不能修改，panic会发生
	setStudent.Field(1).SetInt(19)
	// {joke 19}
	fmt.Println(setStudent)
}

func main() {
	func1()
	func2()
	func3()
}
