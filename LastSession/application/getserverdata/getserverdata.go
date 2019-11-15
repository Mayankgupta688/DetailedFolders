package getserverdata

import (
	"io/ioutil"
	"net/http"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {

	res, _ := http.Get("http://5c055de56b84ee00137d25a0.mockapi.io/api/v1/employeedata")
	data, _ := ioutil.ReadAll(res.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(data))
}
