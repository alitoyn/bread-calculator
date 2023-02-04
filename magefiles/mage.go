//go:build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
)

func Build() error {
	return sh.Run("go", "build", ".")
}

func BuildArm() error {
	return sh.RunWith(map[string]string{"GOARCH": "arm"}, "go", "build", ".")
}

func Test() error {
	output, err := sh.Output("go", "test", "./...")
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}

func Lint() error {
	output, err := sh.Output("golangci-lint", "run", "-E", "goimports")
	if err != nil {
		return err
	}
	fmt.Printf(output)
	return nil
}
