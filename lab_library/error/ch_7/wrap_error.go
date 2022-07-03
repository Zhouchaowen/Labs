package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	return nil, nil
}

func ReadConfig() ([]byte, error) {
	config, err := ReadFile("config.xml")
	if err != nil {
		return config, errors.WithMessage(err, "could not read config")
	}
	return config, nil
}

func main() {
	_, err := ReadConfig()
	if err != nil {
		fmt.Printf("original error; %T  %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n %+v\n", err)
		os.Exit(1)
	}
}
