package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func process() error {
	return errors.Wrap(errors.New("process failed"), "test")
}

func main() {
	fmt.Printf("====== %+v =====\n", process())
	fmt.Printf("===== %v === \n", process())
}
