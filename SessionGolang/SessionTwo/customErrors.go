package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	GetFileData()
}

func GetUserData(age int) (bool, error) {
	if age > 10 {
		return true, nil
	} else {
		return false, errors.New("Wrong Age...")
	}
}

func GetFileData() {
	res, err := os.Open("blankIdentifiers1.go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
