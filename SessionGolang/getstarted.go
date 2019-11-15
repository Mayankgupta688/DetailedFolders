package main

import (
	"fmt"
	"reflect"
)

var Username = "Mayank"
var userCount int

var name, age, designation, salary = "Mayank", 20, "Xyz", 30

func main() {
	floatData := "Hello "
	//intValue := 10
	fmt.Println(reflect.TypeOf(floatData))
	userName := "Mayank"
	userAge := 30
	UpdateValue(&userName, &userAge)
	fmt.Println(userName)
}

func UpdateValue(name *string, age *int) {
	*name = "Anshul"
}
