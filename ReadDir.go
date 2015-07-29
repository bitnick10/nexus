package nexus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FileInfo struct {
	Name  string
	IsDir bool
}

func ReadDir(w http.ResponseWriter, r *http.Request) {
	dirname := r.URL.Query().Get("dirname")
	infos, err := ioutil.ReadDir(dirname)
	myinfos := make([]FileInfo, 0, 10)
	for _, v := range infos {
		newInfo := FileInfo{}
		newInfo.Name = v.Name()
		newInfo.IsDir = v.IsDir()
		myinfos = append(myinfos, newInfo)
	}
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	bol, _ := json.Marshal(myinfos)
	fmt.Fprintf(w, "%s", bol)
	//fmt.Println(string(bol))
}
