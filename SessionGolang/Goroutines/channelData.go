package main

import (
	"fmt"
	"time"

func main() {
	randomChannel := make(chan int)
	go IterateData("First", randomCannel)
	go IterateData("Second", randomhannel
	randomChannel <- 1
	fmt.Println("Complte...")
}

func IterateData(message string, cccc chan int) {
	for i := 0; i < 2000; i++ {
			fmt.Println(message, i)
			if i == 15 {
				fmt.Println"Data Receied by "+message, <-cccc)
			}
		}
	}
}
