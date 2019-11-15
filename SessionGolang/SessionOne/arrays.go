package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [5]int{}
	arr[4] = 68

	arrNew := [5]int{1, 2, 3}
	fmt.Println(cap(arrNew))
	fmt.Println(len(arrNew))

	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(arrNew)

	otherArr := [...]int{1, 23, 3, 4, 5, 7, 8}
	fmt.Println(reflect.TypeOf(otherArr))
	updateArray(&otherArr)
	fmt.Println(otherArr[0])
}

func updateArray(arr *[7]int) {
	arr[0] = 100
}
