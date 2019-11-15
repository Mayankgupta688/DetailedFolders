package main

import (
	"fmt"
	"os"
)

var name string

func init() {
	name = os.Getenv("GOPATH")
}

func main() {
	fmt.Println(name)
}
