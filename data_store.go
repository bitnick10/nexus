package nexus

import (
	// "encoding/json"
	"fmt"
	"net/http"
	// "os"
	"bytes"
	"sync"
	// "strings"
	// "strconv"
)

var nexusList map[string][]string
var lock = sync.RWMutex{}

func init() {
	nexusList = make(map[string][]string)
	// neuxsMap = make(map[string]string)
}
func List(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if r.Method == "POST" {
		lock.Lock()
		defer lock.Unlock()
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		value := buf.String()
		// fmt.Println(key)
		// fmt.Println(value)
		nexusList[key] = append(nexusList[key], value)
		return
	}
	if r.Method == "GET" {
		lock.RLock()
		defer lock.RUnlock()
		if v, ok := nexusList[key]; ok {
			// fmt.Println(len(nexusList[key]))
			// fmt.Println(nexusList[key])
			// fmt.Println(v)
			var buffer bytes.Buffer
			// str := ""
			i := 0
			for i = 0; i < len(v)-1; i++ {
				buffer.WriteString(v[i])
				buffer.WriteString("\n")
				//str += v[i] + "\n"
			}
			if len(v) != 0 {
				buffer.WriteString(v[i])
			}
			fmt.Fprintf(w, "%s", buffer.String())
		} else {
			fmt.Fprintf(w, "%s", "")
		}
		//lock.RUnlock()
	}
	//lock.Unlock()
}

func Map(w http.ResponseWriter, r *http.Request) {
	// key := r.URL.Query().Get("key")
	// if r.Method == "POST" {
	// 	buf := new(bytes.Buffer)
	// 	buf.ReadFrom(r.Body)
	// 	neuxsMap[key] = buf.String()
	// 	return
	// }
	// if r.Method == "GET" {
	// 	if v, ok := neuxsMap[key]; ok {
	// 		fmt.Fprintf(w, "%s", v)
	// 	} else {
	// 		fmt.Fprintf(w, "%s", "")
	// 	}
	// }
}
