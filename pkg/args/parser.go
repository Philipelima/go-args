package args

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

const tagName = "go_arg"


type Parser struct {
	model any
	args  map[string]string
}

func (p *Parser) Parse() error {
	t := reflect.TypeOf(p.model)
	v := reflect.ValueOf(p.model)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if reflect.TypeOf(p.model).Kind() != reflect.Ptr {
		panic("model must be a pointer to a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		tag := field.Tag.Get(tagName)
		if tag == "" {
			continue 
		}

		tagParts := strings.Split(tag, "|")
		argName  := tagParts[0]

		argValue, ok := p.args[argName]
		if !ok || !fieldValue.CanSet() {
			continue
		}

		switch field.Type.Kind() {
			case reflect.String:
				fieldValue.SetString(argValue)
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
				if intVal, err := parseInt(argValue, int(field.Type.Bits())); err == nil {
					fieldValue.SetInt(intVal)
				} else {
					log.Printf("Error converting '%s' to %s: %v\n", argValue, field.Type.Kind(), err)
				}
			default:
				return fmt.Errorf("unsupported type: %s", field.Type.Kind())
		}
	}

	return nil
}


func parseInt(argValue string, bitSize int) (int64, error) {
    return strconv.ParseInt(argValue, 10, bitSize)
}
