package main

import "fmt"

func main() {
	CallDeferFunctions()
}

func CallDeferFunctions() {

	defer CloseConnection()
	defer CloseDBConnection()
	defer CloseConnection()
	defer CloseDBConnection()

	data := 10

	fmt.Println(data)
}

func CloseConnection() {
	fmt.Println("Close Connection")
}

func CloseDBConnection() {
	fmt.Println("Close DB Connection")
}
