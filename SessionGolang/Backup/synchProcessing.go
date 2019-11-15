package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go IterateData("First")
	go IterateData("Second")
	go IterateData("Second1")
	go IterateData("Second2")
	go IterateData("Second3")
	go IterateData("Second4")

	time.Sleep(3 * time.Second)
}

func IterateData(message string) {
	for i := 0; i < 2000; i++ {
		fmt.Println(message, i)
	}
}

func IterateDataAgain() {
	for i := 20; i < 30; i++ {
		fmt.Println("Second", i)
	}
}
