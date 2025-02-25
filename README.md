# go-args


A simple library for parsing command-line arguments in Go.

## Installation

To install the library, run:

```sh
go get github.com/philipelima/go-args/pkg/args
```

## Usage

You can use `go-args` to convert command-line arguments into a map or directly into a struct.

### Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/philipelima/go-args/pkg/args"
)

func main() {
	// show CLI arguments like a map
	argMap := args.AsMap()
	fmt.Println("\nArgs as a Map:", argMap)

	// Define the structure to map the arguments
	var user struct {
		Name string `go_arg:"name"`
		Age  int    `go_arg:"age"`
	}

	// Create a parser and parse the arguments
	parser := args.NewParser(&user)
	if err := parser.Parse(); err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	fmt.Println("\nArgs as a Struct:", user)
}
```

### Running the code

To test, run the following command in the terminal:

```sh
go run example/main.go --name=Jhon --age=30
```

### Expected output:

```
Args as a Map: map[name:Jhon age:30]

Args as a Struct: {Jhon 30}
```

## License

This project is distributed under the MIT license.

