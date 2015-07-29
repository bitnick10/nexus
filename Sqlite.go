package nexus

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func Select1(w http.ResponseWriter, r *http.Request) {
	dbname := r.URL.Query().Get("dbname")
	query := r.URL.Query().Get("query")
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("err at db.Query")
		fmt.Println(err)
	}
	defer rows.Close()
	stringList := make([]string, 0, 10)
	for rows.Next() {
		var str string
		rows.Scan(&str)
		stringList = append(stringList, str)
	}
	bol, _ := json.Marshal(stringList)
	fmt.Fprintf(w, "%s", bol)
}

func Select1Test() {
	dbname := "E:\\z_study\\study.db"
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select id from filter where id like 't%'")
	if err != nil {
		fmt.Println("err at db.Query")
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var str string
		rows.Scan(&str)
		fmt.Println(str)
	}
}
