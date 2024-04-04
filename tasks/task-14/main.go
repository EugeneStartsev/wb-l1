package main

import (
	"fmt"
	"reflect"
)

func main() {
	selectTypeBySwitch([]interface{}{321312, "hello", true, make(chan interface{})})
	selectTypeByReflect([]interface{}{321312, "hello", true, make(chan interface{})})
}

func selectTypeBySwitch(vals []interface{}) {
	for _, val := range vals {
		switch typeVal := val.(type) {
		case string:
			fmt.Printf("this is string %s\n", typeVal)
		case bool:
			fmt.Printf("this is bool %t\n", typeVal)
		case int:
			fmt.Printf("this is int %d\n", typeVal)
		case chan interface{}:
			fmt.Printf("this is channel %v\n", typeVal)
		}
	}
}

func selectTypeByReflect(vals []interface{}) {
	for _, v := range vals {
		res := reflect.ValueOf(v)
		fmt.Printf("%v type is %s\n", res, res.Kind().String())
	}
}
