package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":4000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello All"))
}