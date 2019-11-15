package main

import "fmt"

func main() {
	name, age := GetUserData()
	fmt.Println(name, age)
}

func GetUserData() (name string, age int) {
	return
}
