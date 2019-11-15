package main

import "fmt"

func main() {

	ExecutePanic()
	fmt.Println("Error Recovered....")
}

func ExecutePanic() {
	defer PanicRecovery()
	panic("Data")
}

func PanicRecovery() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
