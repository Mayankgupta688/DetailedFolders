package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", logging(GetUserName))
	http.ListenAndServe(":3000", nil)
}

func GetUserName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello All.. This is the output..."))
}

func ErrorFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Permited"))
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World...")
		if r.FormValue("name") != "Mayank" {
			ErrorFunction(w, r)
		} else {
			f(w, r)
		}

	}
}
