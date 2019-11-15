package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	Name string
	age  string
}

var empList = []Employee{}

func main() {
	http.HandleFunc("/getEmployees", ShowEmployees)
	http.HandleFunc("/addemployee", AddEmployees)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":4000", nil)
}

func ShowEmployees(w http.ResponseWriter, r *http.Request) {
	js, _ := json.Marshal(empList)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func AddEmployees(w http.ResponseWriter, r *http.Request) {

	empData := Employee{r.FormValue("name"), r.FormValue("age")}

	empList = append(empList, Employee{"Ankit", "40"})
	empList = append(empList, empData)

	js, _ := json.Marshal(empList)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {

	employeeList := []Employee{
		{"A", "30"},
		{"B", "30"},
		{"C", "30"},
	}

	employeeList = append(employeeList, Employee{"D", "40"})

	js, err := json.Marshal(employeeList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ShowOtherBooks(w http.ResponseWriter, r *http.Request) {

	employeeList := []Employee{
		{"A", "30"},
		{"B", "30"},
		{"C", "30"},
	}

	employeeList = append(employeeList, Employee{"D", "40"})

	js, err := json.Marshal(employeeList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(js)
}
