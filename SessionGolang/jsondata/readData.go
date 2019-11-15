package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Employees struct {
	Employees [90]Employee `json:"employees"`
}

type Employee struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func main() {
	jsonFile, _ := ioutil.ReadFile("../data.json")

	var employees Employees
	json.Unmarshal(jsonFile, &employees)

	for i := 0; i < len(employees.Employees); i++ {
		fmt.Println("User Type: " + employees.Employees[i].Name)
	}

	fmt.Println(len(employees.Employees))
}
