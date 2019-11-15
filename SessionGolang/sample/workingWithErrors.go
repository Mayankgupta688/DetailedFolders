package main

import (
	"errors"
	"fmt"
)

func main() {
	val, err := GetCustomError()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}

func GetCustomError() (int, error) {
	return 0, errors.New("Invalid")
}
