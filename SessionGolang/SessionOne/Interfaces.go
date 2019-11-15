package main

import "fmt"

type Person interface {
	UpdateDetails()
	ShowName()
}

type Employee struct {
	Name string
}

func (emp *Employee) UpdateDetails() {
	fmt.Println(emp)
}

func (emp Employee) ShowName() {
	emp.Name = "Updated Employee"
}

type Manager struct {
	Name string
}

func (emp *Manager) UpdateDetails() {
	emp.Name = "Updated Employee"
}

func (emp Manager) ShowName() {
	fmt.Println(emp.Name)
}

func main() {

	var sampleInterface Person

	var employeeSimple Employee

	employeeSimple = Employee{}

	employeeSimple.Name = "Mayank"
	sampleInterface = employeeSimple
	sampleInterface.UpdateDetails()
	sampleInterface.ShowName()

	managerSimple := new(Manager)
	managerSimple.Name = "Manager"
	sampleInterface = managerSimple
	sampleInterface.UpdateDetails()
	sampleInterface.ShowName()

	fmt.Println("Data", &sampleInterface)
	fmt.Println(&employeeSimple)
}
