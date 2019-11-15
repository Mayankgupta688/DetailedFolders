package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, _ := ioutil.ReadFile("sample.txt")
	fmt.Println(string(data))

	mydata := []byte("all my data I want to write to a file")
	ioutil.WriteFile("new.txt", mydata, 0777)

	f, _ := os.OpenFile("new.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if _, err := f.WriteString("new data that wasn't there originally\n"); err != nil {
		panic(err)
	}

	defer f.Close()
}
