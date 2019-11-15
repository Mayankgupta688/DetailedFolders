package main

import (
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/getuser", HelloServer)
	http.ListenAndServe(":4000", nil)
}

type Employee struct {
	Name string
	Age string
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	emp := []Employee{{"Name", "30"}, {"Name", "30"}, {"Name", "30"}}
	jsonData, _ := json.Marshal(emp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}