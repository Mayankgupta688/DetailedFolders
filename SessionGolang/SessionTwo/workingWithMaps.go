package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	stud := Student{"M"}
	otherMap := make(map[Student]string)
	otherMap[stud] = "Mayank"

	fmt.Println(otherMap[stud])

	for key, _ := range otherMap {
		fmt.Println(key)
	}
	sampleMap := map[string]string{
		"name": "Mayank",
		"age1": "56",
	}

	UpdateData(sampleMap)

	fmt.Println(sampleMap["name"])
}

func UpdateData(mapData map[string]string) {
	mapData["name"] = "dhfg"
}
