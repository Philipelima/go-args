package args

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const tag_name = "go_arg"

type Parser struct {
	model interface{}
	args  map[string]string
}

func (p *Parser) Parse() {
	t := reflect.TypeOf(p.model)
	v := reflect.ValueOf(p.model)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		tag := field.Tag.Get(tag_name)
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
			case reflect.Int, reflect.Int64:
				if intVal, err := strconv.ParseInt(argValue, 10, 64); err == nil {
					fieldValue.SetInt(intVal)
				} else {
					fmt.Printf("Erro ao converter '%s' para int64: %v\n", argValue, err)
				}
			case  reflect.Int32:
				if intVal, err := strconv.ParseInt(argValue, 10, 34); err == nil {
					fieldValue.SetInt(intVal)
				} else {
					fmt.Printf("Erro ao converter '%s' para int34: %v\n", argValue, err)
				}
			case reflect.Int16:
				if intVal, err := strconv.ParseInt(argValue, 10, 16); err == nil {
					fieldValue.SetInt(intVal)
				} else {
					fmt.Printf("Erro ao converter '%s' para int16: %v\n", argValue, err)
				}
			case reflect.Int8:
				if intVal, err := strconv.ParseInt(argValue, 10, 8); err == nil {
					fieldValue.SetInt(intVal)
				} else {
					fmt.Printf("Erro ao converter '%s' para int8: %v\n", argValue, err)
				}
			default:
				fmt.Printf("Tipo não suportado: %s\n", field.Type.Kind())
		}
	}
}

