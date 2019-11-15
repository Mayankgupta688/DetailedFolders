package main

import "fmt"

type Address struct {
	Street string
	State  string
}

type Employee struct {
	Name        string
	Age         int
	Designation string
	location    Address
}

func (emp *Employee) ShowAge() {
	emp.Age = 50
}

func main() {
	emp := Employee{"Mayank", 10, "Xyz", Address{"Delhi", "Y"}}
	emp.ShowAge()
	fmt.Println(emp.location.Street)
}
