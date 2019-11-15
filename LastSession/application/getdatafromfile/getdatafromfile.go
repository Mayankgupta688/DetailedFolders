package getdatafromfile

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetDataFromJsonFile(w http.ResponseWriter, r *http.Request) {
	res, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}
