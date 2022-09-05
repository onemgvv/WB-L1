package main

import (
	"fmt"
	"reflect"
)

func typeDef(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func typeReflect(v interface{}) reflect.Kind {
	return reflect.TypeOf(v).Kind()
}

func typeSwitch(v interface{}) string {
	switch tp := v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return fmt.Sprintf("%T", tp)
	}
}

func main() {
	chInt := make(chan int)
	chStr := make(chan string)
	chBool := make(chan bool)

	fmt.Println("Type by typeDef:", typeDef(true))
	fmt.Println("Type by typeDef:", typeDef(chInt))
	fmt.Println("Type by reflect:", typeReflect(8))
	fmt.Println("Type by reflect:", typeReflect(chStr))
	fmt.Println("Type by switch:", typeSwitch("she"))
	fmt.Println("Type by switch:", typeSwitch(chBool))
}