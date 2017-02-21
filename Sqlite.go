package nexus

// cygwin go-sqlite3 build gcc http://tdm-gcc.tdragon.net/

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"sync"
	// "strings"
	// "strconv"
)

var cache map[string]string // := make(map[string]string)
var lock = sync.RWMutex{}
var neuxsMap map[string]string

func init() {
	cache = make(map[string]string)
	neuxsMap = make(map[string]string)
}
func Map(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if r.Method == "POST" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		neuxsMap[key] = buf.String()
		return
	}
	if r.Method == "GET" {
		if v, ok := neuxsMap[key]; ok {
			fmt.Fprintf(w, "%s", v)
		} else {
			fmt.Fprintf(w, "%s", "")
		}
	}
}
func Select(w http.ResponseWriter, r *http.Request) (string, error) {
	dbname := r.URL.Query().Get("dbname")
	query := r.URL.Query().Get("query")
	lock.RLock()
	if v, ok := cache[r.URL.String()]; ok {
		fmt.Fprintf(w, "%s", v)
		fmt.Println("cache")
		lock.RUnlock()
		return "", nil
	}
	lock.RUnlock()
	if _, err := os.Stat(dbname); os.IsNotExist(err) {
		// database does not exist
		fmt.Println(dbname + " database does not exist")
		fmt.Fprintf(w, "%s", "database does not exist")
		return "", nil
	}
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	// d := json.NewDecoder(strings.NewReader(string(tableData)))
	// d.UseNumber()
	// d.Decode()
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	// fmt.Println(string(jsonData))
	lock.Lock()
	cache[r.URL.String()] = string(jsonData)
	lock.Unlock()
	fmt.Fprintf(w, "%s", string(jsonData))
	return string(jsonData), nil
}

func Select1(w http.ResponseWriter, r *http.Request) {
	dbname := r.URL.Query().Get("dbname")
	query := r.URL.Query().Get("query")
	fmt.Println("dbname ", dbname)
	fmt.Println("query ", query)
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

func Select7(w http.ResponseWriter, r *http.Request) {
	dbname := r.URL.Query().Get("dbname")
	query := r.URL.Query().Get("query")
	fmt.Println("dbname ", dbname)
	fmt.Println("query ", query)
	db, err := sql.Open("sqlite3", dbname)
	//fmt.Println("Open after ")
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	// fmt.Println("Open()")
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("err at db.Query")
		fmt.Println(err)
	}
	// fmt.Println("Query()")
	defer rows.Close()
	stringList := make([][]string, 0, 10)
	// fmt.Println("for rows.Next()")
	for rows.Next() {
		var str1 int
		var str2 float64
		var str3 float64
		var str4 float64
		var str5 float64
		var str6 float64
		var str7 float64
		err = rows.Scan(&str1, &str2, &str3, &str4, &str5, &str6, &str7)
		if err != nil {
			fmt.Println("err rows.Scan")
			fmt.Println(err)
		}
		// fmt.Println(str1,str2)
		// fmt.Println(str2)
		onedaydate := make([]string, 0, 7)
		onedaydate = append(onedaydate, fmt.Sprintf("%d", str1))
		onedaydate = append(onedaydate, fmt.Sprintf("%.2f", str2))
		onedaydate = append(onedaydate, fmt.Sprintf("%.2f", str3))
		onedaydate = append(onedaydate, fmt.Sprintf("%.2f", str4))
		onedaydate = append(onedaydate, fmt.Sprintf("%.2f", str5))
		onedaydate = append(onedaydate, fmt.Sprintf("%.2f", str6))
		onedaydate = append(onedaydate, fmt.Sprintf("%.2f", str7))
		stringList = append(stringList, onedaydate)
	}
	//fmt.Println("for rows.Next() end")
	bol, _ := json.Marshal(stringList)
	fmt.Fprintf(w, "%s", bol)
}

// func Select1Test() {
// 	dbname := "E:\\z_study\\study.db"
// 	db, err := sql.Open("sqlite3", dbname)
// 	if err != nil {
// 		fmt.Println("err")
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("select id from filter where id like 't%'")
// 	if err != nil {
// 		fmt.Println("err at db.Query")
// 		fmt.Println(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var str string
// 		rows.Scan(&str)
// 		fmt.Println(str)
// 	}
// }
