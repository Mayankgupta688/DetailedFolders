package main

import "fmt"

func main() {
	var a, b, c = GetMultipleValues()
	fmt.Println(a, b, c)
}

func GetMultipleValues() (int, string, bool) {
	return 1, "Mayank", false
}
