package main

import (
	"fmt"
	"runtime"
	"sync"
)

var syncLock sync.WaitGroup
var mutex = &sync.Mutex{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	syncLock.Add(3)
	randomChannel := make(chan string, 3)
	go IterateData("First", randomChannel)
	go IterateData("Second", randomChannel)
	go IterateData("Third", randomChannel)
	syncLock.Wait()

	a := <-randomChannel
	b := <-randomChannel
	c := <-randomChannel

	fmt.Println(a, b, c)
}

func RandomCalled() {
	syncLock.Done()
}

func IterateData(message string, cccc chan string) {
	// mutex.Lock()

	cccc <- message
	// mutex.Unlock()
	go RandomCalled()
}
