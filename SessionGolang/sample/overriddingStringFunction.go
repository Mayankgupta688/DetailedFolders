package main

import "fmt"

type Employee struct {
	name string
}

func (emp Employee) String() string {
	return "Hello"
}

type money int

func main() {
	emp := Employee{"mayank"}
	fmt.Println(emp.String())
}
