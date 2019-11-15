package main

import "fmt"

type Manager struct {
	Name     string
	TeamSize int
	Employee
	Person
}

type Employee struct {
	Name string
	Age  int
}

type Person struct {
	UserName string
	UserAge  int
}

func (emp Person) ShowDetails() int {
	return emp.UserAge
}

func (emp Employee) ShowDetails() int {
	return emp.Age
}

func (emp Manager) ShowDetails() {
	fmt.Println(emp.Person.ShowDetails() + emp.Employee.ShowDetails())

}

func main() {
	manag := Manager{}
	manag.Name = "Mayank"
	manag.Age = 50
	manag.Employee.Name = "Anshul"
	manag.TeamSize = 50
	fmt.Println(manag.Employee.Name)
	fmt.Println(manag)
	manag.UserName = "sdgysgd"
	manag.UserAge = 78
	manag.ShowDetails()
}
