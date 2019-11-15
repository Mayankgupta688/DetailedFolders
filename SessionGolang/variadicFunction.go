package main

import (
	"fmt"
	"reflect"
)

var name string

func main() {
	var name string
	name = "Anshul"
	VariadicFunction("Mayank", 1, 2, 3, 4)
	VariadicFunction("Anshul", 1, 2, 3, 4, 5, 6)
	fmt.Println(name)
}

func VariadicFunction(data string, inputData ...int) {
	fmt.Println(cap(inputData))
	fmt.Println(reflect.TypeOf(inputData))
}
