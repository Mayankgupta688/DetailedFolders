package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleIndex)
	http.ListenAndServe(":3000", nil)
}

type Employee struct {
	name string
	age  int
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	//newEmp := Employee{"Mayank", 30}
	w.Write([]byte("Hello World"))

}
