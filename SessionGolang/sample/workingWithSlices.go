package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5}
	sampleSlice[2] = 30
	UpdateSlice(sampleSlice)
	sampleSlice = append(sampleSlice, 10)
	fmt.Println(sampleSlice[4])
	fmt.Println(cap(sampleSlice))
}

func UpdateSlice(data []int) {
	data[4] = 40
}
