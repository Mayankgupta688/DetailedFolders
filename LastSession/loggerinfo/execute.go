package main

import (
	"io"
	"log"
	"os"
)

func main() {
	Panic("Error.... Hello")
}

func Panic(logString string) {
	f, err := os.OpenFile("./tmp/orders.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	//log.Fatal(logString)
	log.Fatalln(logString)
}
