package main

import (
	"Labs/lab_library/validator/check"
	"fmt"
)

var checker *check.Check

func init() {
	checker = check.NewCheck()
}

// 校验map
func main() {
	// 数据
	data := map[string]interface{}{"name": "", "email": ""}
	// 校验规则
	// name:必填，最小长度8，最大长度32，email:为空时忽略校验，不为空时必须是邮件格式
	rules := map[string]interface{}{"name": "required,min=2,max=5", "email": "omitempty,email"}
	err := checker.ValidateMap(data, rules)
	fmt.Println(err)

	data = map[string]interface{}{"name": "1", "email": ""}
	err = checker.ValidateMap(data, rules)
	fmt.Println(err)

	data = map[string]interface{}{"name": "123456", "email": "1"}
	err = checker.ValidateMap(data, rules)
	fmt.Println(err)

	data = map[string]interface{}{"name": "123", "email": "123456@qq.com"}
	err = checker.ValidateMap(data, rules)
	fmt.Println(err)
}
