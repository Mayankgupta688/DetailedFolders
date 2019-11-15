package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Name        string
	Age         int
	Designation string
}

func main() {
	newEmp := Employee{"Mayank", 20, "sdsydg"}
	fmt.Println(newEmp.Name)
	fmt.Println(newEmp.Age)

	otherEmployee := Employee{}
	otherEmployee.Name = "Mayank Gupta"
	fmt.Println(otherEmployee.Name)

	var newData Employee
	fmt.Println("Age ", newData)

	sampleEmployee := Employee{Age: 50, Name: "Mayank"}
	sampleEmployee.Designation = "shjdf"
	fmt.Println((sampleEmployee.Name))

	fmt.Println(reflect.TypeOf(sampleEmployee))
}
