package main

import (
	"SessionGolang/sample/infopkg"
	"fmt"
)

func init() {
	fmt.Println("Parent")
}

func main() {
	infopkg.GetUserName()

	newStud := new(infopkg.Student)
	newStud.Age = 10
	fmt.Println(infopkg.StudentData.Age)
}
