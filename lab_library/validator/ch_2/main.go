package main

import (
	"Labs/lab_library/validator/check"
	"fmt"
)

var checker *check.Check

func init() {
	checker = check.NewCheck()
}

func main() {
	var array = []string{"1", ""}
	err := checker.Var(
		array,
		"required,dive,required",
	)
	fmt.Println(err)

	var number int = 10
	err = checker.Var(
		number,
		"required,len=9",
	)
	fmt.Println(err)

	m := map[string]string{"val1": "val1", "val2": "val2", "val3": "val3"}
	err = checker.Var(m, "required,dive,keys,required,endkeys,required")
	fmt.Println(err)
}
