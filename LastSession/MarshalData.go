package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetUserData)
	http.ListenAndServe(":4000", nil)
}

type Employee struct {
	Id        string `json:"id"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
}

type Employees struct {
	Employees [5]Employee `json:"employees"`
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://5c055de56b84ee00137d25a0.mockapi.io/api/v1/employeedata")
	data, _ := ioutil.ReadAll(res.Body)
	var empList Employees
	json.Unmarshal(data, &empList)
	fmt.Println(len(empList.Employees))
	jsonData, _ := json.Marshal(empList.Employees)
	w.Write(jsonData)
}
