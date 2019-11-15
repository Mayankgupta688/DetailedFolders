package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := httprouter.New()
	r.GET("/", GetIndexPage)
	r.POST("/addtable", AddTable)
	r.POST("/adddatatotable", AddDataToTable)
	r.GET("/getsingleusers/:id", GetSingleUsers)
	r.GET("/getallusers", GetAllUsers)
	r.POST("/deleteuserwithid/2", DeleteUserWithId)
	http.ListenAndServe(":3000", r)
}

func AddTable(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, _ := sql.Open("mysql", "7N10fxmDYp:NnktEISo8u@(remotemysql.com:3306)/7N10fxmDYp?parseTime=true")

	query := `
    CREATE TABLE newstudent (
        id TEXT NOT NULL,
        createdAt TEXT NOT NULL,
        name TEXT NOT NULL,
        avatar TEXT
	);`

	_, err := db.Exec(query)

	fmt.Println(err)

	rw.Write([]byte("Database Added Successfully"))

	db.Close()
}

func AddDataToTable(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, _ := sql.Open("mysql", "7N10fxmDYp:NnktEISo8u@(remotemysql.com:3306)/7N10fxmDYp?parseTime=true")

	_, err := db.Exec(`INSERT INTO newstudent (id, createdAt, name, avatar) VALUES (?, ?, ?, ?)`,
		r.FormValue("id"), r.FormValue("createdAt"), r.FormValue("name"), r.FormValue("avatar"))

	fmt.Println(err)

	fp := path.Join("public", "index.html")
	tmpl, _ := template.ParseFiles(fp)
	tmpl.Execute(rw, nil)

	db.Close()
}

func DeleteUserWithId(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, _ := sql.Open("mysql", "7N10fxmDYp:NnktEISo8u@(remotemysql.com:3306)/7N10fxmDYp?parseTime=true")
	//DELETE FROM `newstudent` WHERE id=1

	fmt.Println("Delete Called")
	_, err := db.Exec(`DELETE FROM newstudent WHERE id=?`, r.FormValue("id"))

	fmt.Println(err)

	db.Close()
	w.Write([]byte("Done..."))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, _ := sql.Open("mysql", "7N10fxmDYp:NnktEISo8u@(remotemysql.com:3306)/7N10fxmDYp?parseTime=true")
	rows, _ := db.Query(`SELECT id, createdAt, name, avatar FROM newstudent`) // check err
	defer rows.Close()

	type User struct {
		Id        string
		CreatedAt string
		Name      string
		Avatar    string
	}

	var users []User
	for rows.Next() {
		var u User
		_ = rows.Scan(&u.Id, &u.CreatedAt, &u.Name, &u.Avatar) // check err
		users = append(users, u)
	}

	// jsonData, _ := json.Marshal(users)
	// //jsonData := string(data)
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Write(jsonData)

	fp := path.Join("public", "list.html")
	tmpl, _ := template.ParseFiles(fp)
	tmpl.Execute(w, users)

}

func GetSingleUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, _ := sql.Open("mysql", "7N10fxmDYp:NnktEISo8u@(remotemysql.com:3306)/7N10fxmDYp?parseTime=true")

	type User struct {
		Id        string
		CreatedAt string
		Name      string
		Avatar    string
	}

	var u User
	query := "SELECT id, createdAt, name, avatar FROM newstudent WHERE id =" + p.ByName("id")
	_ = db.QueryRow(query).Scan(&u.Id, &u.CreatedAt, &u.Name, &u.Avatar)

	// jsonData, _ := json.Marshal(users)
	// //jsonData := string(data)
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Write(jsonData)

	fp := path.Join("public", "list.html")
	tmpl, _ := template.ParseFiles(fp)
	fmt.Println(u)
	tmpl.Execute(w, u)

}

func GetIndexPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fp := path.Join("public", "index.html")
	tmpl, _ := template.ParseFiles(fp)
	tmpl.Execute(w, nil)
}
