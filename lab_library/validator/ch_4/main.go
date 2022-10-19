package main

import (
	"Labs/lab_library/validator/check"
	"fmt"
)

type Struct2 struct {
	// 不能为空，长度大于0，元素不能有空值
	Array []string `validate:"required,gt=0,dive,required"`
}
type Struct3 struct {
	// 不能为空，长度大于0，key值长度最大为5，value值不能为空并且最大长度为2
	Map map[string]string `validate:"required,gt=0,dive,keys,keymax,endkeys,required,max=2"`
}

type Struct4 struct {
	// 不能为空，内容不能存在空值
	Column1 []*Struct5 `validate:"required,dive,required"`
}

type Struct5 struct {
	// 不能为空
	Column2 string `validate:"required"`
}

var checker *check.Check

func init() {
	checker = check.NewCheck()
}

func main() {
	var test Struct2
	// 不能为空
	err := checker.Struct(test)
	fmt.Println(err)

	// 长度大于0
	test.Array = []string{}
	err = checker.Struct(test)
	fmt.Println(err)
	// 元素不能为空
	test.Array = []string{"11", ""}
	err = checker.Struct(test)
	fmt.Println(err)

	var test1 Struct3
	checker.RegisterAlias("keymax", "max=5")

	// 不能为空
	err = checker.Struct(test1)
	fmt.Println(err)
	// 长度大于0
	test1.Map = map[string]string{}
	err = checker.Struct(test1)
	fmt.Println(err)

	// key值长度最大为5，value值不能为空
	test1.Map = map[string]string{"123456": "1", "12345": ""}
	err = checker.Struct(test1)
	fmt.Println(err)

	test1.Map = map[string]string{"123": "1", "1234": "2", "12345": "333"}
	err = checker.Struct(test1)
	fmt.Println(err)

	var test2 Struct4
	err = checker.Struct(test2)
	fmt.Println(err)

	test2.Column1 = []*Struct5{{Column2: ""}, nil}
	err = checker.Struct(test2)
	fmt.Println(err)
}
