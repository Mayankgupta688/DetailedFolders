package main

import "net/http"

func main() {
	http.HandleFunc("/", HelloFunc)
	http.ListenAndServe(":3000", nil)
}

func HelloFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
