package main

import "fmt"

type Employee struct {
	Name string
}

func main() {
	emp := Employee{"M"}
	UpdateStruct(&emp)
	fmt.Println(emp.Name)
}

func UpdateStruct(employee *Employee) {
	employee.Name = "XYZ"
}
