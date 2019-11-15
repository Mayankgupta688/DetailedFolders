package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)
func main() {
	response, _ := http.Get("http://5c055de56b84ee00137d25a0.mockapi.io/api/v1/employees")
	responseData, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseData))
}