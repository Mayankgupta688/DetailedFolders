package main

import (
	"fmt"
	"reflect"
)

func main() {
	userMap := map[string]string{
		"user":        "Mayank",
		"age":         "20",
		"designation": "Developer",
	}

	otherUserMap := make(map[string]string)
	otherUserMap["name"] = "Mayank"

	fmt.Println(reflect.TypeOf(otherUserMap))
	fmt.Println(reflect.TypeOf(userMap))

	UpdateMapData(&userMap)
	fmt.Println("Value ", userMap["name"])
}

func UpdateMapData(data *map[string]string) {
	(*data)["name"] = "AAA"
}
