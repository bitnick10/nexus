package nexus

import (
	// "bytes"
	"database/sql"
	// "encoding/json"
	"env"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	// "net/url"
	// "os"
	// "sync"
	// "strings"
	// "strconv"
)

func MySQLSelect(w http.ResponseWriter, r *http.Request) (string, error) {
	dbName := r.URL.Query().Get("dbname") // case sensitive
	query := r.URL.Query().Get("query")
	// query, _ = url.QueryUnescape(query)
	// fmt.Println("quey a a  a ", query)
	//connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, "postgres", password, dbname)
	// fmt.Println(connectString)
	db, err := sql.Open("mysql", env.MySQLConnectionString("tcp(192.168.0.170:3306)", dbName))
	// fmt.Println(dbName)
	defer db.Close()
	if err != nil {
		fmt.Println("mysql conncet err")
		panic(err)
	}
	// fmt.Println("111111111111111111111111111")
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("mysql query err")
		fmt.Println(err)
	}
	// fmt.Println("222222222222222222222222")
	defer rows.Close()
	js := RowsToJSONString(rows)
	// fmt.Println("33333333333333333333333")
	fmt.Fprintf(w, "%s", js)
	return "", nil
}

func MySQLInsert(w http.ResponseWriter, r *http.Request) (string, error) {
	dbName := r.URL.Query().Get("dbname") // case sensitive
	query := r.URL.Query().Get("query")
	// query, _ = url.QueryUnescape(query)
	// fmt.Println("quey a a  a ", query)
	//connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, "postgres", password, dbname)
	// fmt.Println(connectString)
	db, err := sql.Open("mysql", env.MySQLConnectionString("tcp(192.168.0.170:3306)", dbName))
	// fmt.Println(dbName)
	defer db.Close()
	if err != nil {
		fmt.Println("mysql conncet err")
		panic(err)
	}
	//fmt.Println(query)
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println("query error", query)
		fmt.Println(err)
		return "", err
	}
	lastInsertId, _ := res.LastInsertId()
	fmt.Fprintf(w, "%d", lastInsertId)
	return "", nil
}
