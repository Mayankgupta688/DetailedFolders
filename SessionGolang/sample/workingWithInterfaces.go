package main

import "fmt"

type EmployeeInt interface {
	ShowUserData()
}

type TeamLead struct {
	name        string
	age         int
	designation string
}

func (lead TeamLead) ShowUserData() {
	fmt.Println("Lead")
}

type TeamManager struct {
	name        string
	age         int
	designation string
}

func (lead TeamManager) ShowUserData() {
	fmt.Println("Manager")
}

var empInterface EmployeeInt

func main() {

	newManager := TeamManager{"Mayank", 12, "Developer"}
	empInterface = newManager
	empInterface.ShowUserData()
}
