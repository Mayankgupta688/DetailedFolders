package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	arr := []Student{{"M", 10}, {"A", 10}, {"Q", 10}}
	UpdateData(arr)
	fmt.Println(&arr[0])
	//fmt.Println(arrOne[0])
	// fmt.Println(len(arr))
	// fmt.Println(cap(arr))
}

func UpdateData(sliceData []Student) {
	sliceData[0].Age = 1000
}
