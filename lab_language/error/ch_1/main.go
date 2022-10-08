// errors.New()返回指针，不能直接用于错误相等
package main

import (
	"errors"
	"fmt"
)

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

/*
自定义Err的坑
func New(text string) error {
	return &errorString{text}  返回指针
}
*/
func main() {
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}

	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}
}
