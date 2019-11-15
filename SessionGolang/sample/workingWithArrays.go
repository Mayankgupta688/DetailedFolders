package main

import "fmt"

func main() {
	userArray := [8]int{1, 2}

	dynamicArray := [...]int{1, 2}

	UpdateArray(&userArray)
	fmt.Println(userArray[0])
	fmt.Println(len(dynamicArray))
}

func UpdateArray(arr *[8]int) {
	arr[0] = 12
}
