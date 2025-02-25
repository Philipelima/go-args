package main

import (
	"fmt"
	"log"

	"github.com/philipelima/go-args/pkg/args"
)

func main() {

	argMap := args.AsMap()
	fmt.Println("\nArgs as a Map:", argMap)

	var user struct {
		Name string `go_arg:"name"`
		Age  int    `go_arg:"age"`
	}

	parser := args.NewParser(&user)
	if err := parser.Parse(); err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	fmt.Println("\nArgs as a Struct:", user)
}
