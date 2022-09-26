package main

import "fmt"

type Name struct {
	A string
}

func New(a int) *Name {
	if a == 0 {
		return &Name{"a"}
	}
	return nil
}

func (a *Name) get() error {
	fmt.Println("ok", a.A)
	return nil
}
func main() {
	a := New(1)
	defer a.get()

	fmt.Println("success")
}
