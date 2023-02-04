//go:build mage

package main

import (
	"fmt"
    "github.com/magefile/mage/sh"
)

func Build() error {
    return sh.Run("go", "build", ".")
}

func Test() error {
    output, err := sh.Output("go", "test", "./...")
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
