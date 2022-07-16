package main

import "fmt"

type Option func(opts *Options)

func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

type Options struct {
	A string
	B string
	C string
	D string
	E string
}

func SetA(a string) Option {
	return func(ops *Options) {
		ops.A = a
	}
}

func SetB(b string) Option {
	return func(ops *Options) {
		ops.B = b
	}
}

func SetC(c string) Option {
	return func(ops *Options) {
		ops.C = c
	}
}

func SetD(d string) Option {
	return func(ops *Options) {
		ops.D = d
	}
}

func main() {
	op := loadOptions(SetA("a"), SetB("b"), SetC("c"), SetD("d"))

	fmt.Println(op)
}
