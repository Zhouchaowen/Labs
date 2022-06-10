package main

import (
	"fmt"
	"reflect"
)

func func1() {
	var a int = 50
	v := reflect.ValueOf(a) // 返回Value类型对象，值为50
	t := reflect.TypeOf(a)  // 返回Type类型对象，值为int
	//50 int int int
	fmt.Println(v, t, v.Type(), t.Kind())
}

func func2() {
	var b [5]int = [5]int{5, 6, 7, 8}
	// [5]int array int
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem())
}

type age interface {
	Age()
}

type sex interface {
	Sex()
}

type Student struct {
	name string
	age  int
}

func (s Student) Name() {
	fmt.Println(s.name)
}

func (s Student) Age() {
	fmt.Println(s.age)
}

func func3() {
	var Pupil Student
	// main.Student
	p := reflect.ValueOf(Pupil) // 使用ValueOf()获取到结构体的Value对象
	// struct
	fmt.Println("Type: ", p.Type()) // 输出:Student
	fmt.Println("Kind: ", p.Kind()) // 输出:struct

	t := p.Type()

	fmt.Println("Align: ", t.Align())
	fmt.Println("FieldAlign: ", t.FieldAlign())
	fmt.Println("Method: ", t.Method(0))
	m, b := t.MethodByName("Name")
	fmt.Println("MethodByName: ", m, b)
	fmt.Println("NumMethod: ", t.NumMethod())
	fmt.Println("Name: ", t.Name())
	fmt.Println("PkgPath: ", t.PkgPath())
	fmt.Println("Size: ", t.Size())
	fmt.Println("String: ", t.String())
	fmt.Println("Kind: ", t.Kind())
	fmt.Println("Implements: ", t.Implements(reflect.TypeOf(new(age)).Elem()))
	fmt.Println("AssignableTo: ", t.AssignableTo(reflect.TypeOf(new(sex)).Elem()))
	fmt.Println("ConvertibleTo: ", t.ConvertibleTo(reflect.TypeOf(new(age)).Elem()))
	fmt.Println("Comparable: ", t.Comparable())
	//fmt.Println("Bits: ",t.Bits())
	//fmt.Println("ChanDir: ",t.ChanDir())
	//fmt.Println("IsVariadic: ",t.IsVariadic())
	//fmt.Println("Elem: ",t.Elem())
	fmt.Println("Field: ", t.Field(0))
	//fmt.Println("Comparable: ",t.FieldByIndex([]int{1,2}))
	mm, b := t.FieldByName("Name")
	fmt.Println("FieldByName: ", mm, b)
	//fmt.Println("FieldByNameFunc: ",t.FieldByNameFunc())
	//fmt.Println("In: ",t.In(0))
	//fmt.Println("Key: ",t.Key())
	//fmt.Println("Len: ",t.Len())
	fmt.Println("NumField: ", t.NumField())
	//fmt.Println("NumIn: ",t.NumIn())
	//fmt.Println("NumOut: ",t.NumOut())
	//fmt.Println("Out: ",t.Out(0))
}

func main() {
	//func1()
	//func2()
	func3()
}
