package main

import (
	"Labs/lab_library/validator/check"
	"fmt"
)

type StructCustom struct {
	ULen         string `validate:"ulen=2"`
	UGreaterThan string `validate:"ugt=2"`
	ULessThan    string `validate:"ult=2"`
}

var checker *check.Check

func init() {
	checker = check.NewCheck()
}
func main() {
	var test StructCustom
	test.ULen = "张"
	err := checker.Struct(test)
	fmt.Println(err)

	test.ULen = "张三"
	test.UGreaterThan = "李四"
	err = checker.Struct(test)
	fmt.Println(err)

	test.ULen = "张三"
	test.UGreaterThan = "李四啊"
	test.ULessThan = "王五"
	err = checker.Struct(test)
	fmt.Println(err)

}
