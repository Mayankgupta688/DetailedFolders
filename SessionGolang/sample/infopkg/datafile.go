package infopkg

import "fmt"

func init() {
	fmt.Println("Child")
}

func GetUserName() {
	fmt.Println("User Data")
}

type Student struct {
	name string
	Age  int
}

var StudentData = Student{"Mayank", 30}
