package main

import (
	"fmt"

	"github.com/philipelima/go-args/pkg/args"
)

func main() {

	ars := args.AsMap()

	fmt.Println("\nArgs as a Map: ", ars)

	var User struct {
		Name string `go_arg:"name|required"`
		Age  int64  `go_arg:"age"`
	}

	parser := args.NewParser(&User)
	parser.Parse()

	fmt.Println("\nArgs as a Struct: ", User)
}
