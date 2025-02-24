package main

import (
	"fmt"

	"github.com/philipelima/go-args/pkg/args"
)

func main() {

	ars := args.AsMap()
	
	fmt.Println("\nArgs as a Map: ", ars)

	
	var Arguments struct {
		Hello string `go_arg:"name:hello"`
		Size  int64  `go_arg:"name:size"`
	}

	parser := args.NewParser(&Arguments)
	parser.Parse()


	fmt.Println("\nArgs as a Struct: ", Arguments.Size)
}
