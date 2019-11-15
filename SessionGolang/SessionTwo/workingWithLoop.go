package main

import (
	"fmt"
	"os"
	_ "reflect"
)

func main() {
	res, _ := os.OpenFile("sample.txt")
	fmt.Println(res)
	arr := [8]int{10, 20, 30, 4, 5, 6, 7, 8}

	for _, value := range arr {
		fmt.Println(value)
	}

	for i, j := 0, 10; i < 10; i++ {
		fmt.Println("Value of: ", i, j)
	}

	fmt.Println("Value of: ", i)
}
