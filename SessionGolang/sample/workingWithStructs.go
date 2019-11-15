package main

import "fmt"

type Employee struct {
	name string
	age  int
	add  [2]Address
}

type Address struct {
	street string
	state  string
}

func main() {

	newEmp := new(Employee)
	newEmp.add = [2]Address{{"A", "B"}, {"A", "B"}}

	fmt.Println(len(newEmp.add))
}
