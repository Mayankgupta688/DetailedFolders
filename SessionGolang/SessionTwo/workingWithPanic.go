package main

import "fmt"

func main() {

	UpdateCode("XYZ", 0)

	fmt.Println("Function Recovered....")
}

func UpdateCode(name string, val int) {
	defer RecoverFunction()
	if name == "XYZ" {
		(func() {
			defer RecoverFunction()
			fmt.Println("This is Executing...")
			panic("sfsdjdfsu")
		})()
		fmt.Println("This is Still... Executing...")
		fmt.Println("Done")
	} else {
		fmt.Println("Done")
	}

}

func RecoverFunction() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("I can Recover...")
}
