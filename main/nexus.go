package main

import (
	"../../nexus"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	go server()
	var input string
	fmt.Scanln(&input)
	fmt.Println("server closed")
}
func server() {
	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Println(r.URL)
		fmt.Println(r.Method)
		//fmt.Println(r.Body)
		//buf := new(bytes.Buffer)
		//buf.ReadFrom(r.Body)
		//s := buf.String()
		//fmt.Println(s)
		str := r.URL.Query().Get("content")
		fmt.Println(str)
		fmt.Fprintf(w, "%s", str+" too")
	})
	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		nexus.Map(w, r)
	})
	http.HandleFunc("/command", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		request, _ := url.QueryUnescape(r.URL.RequestURI())
		fmt.Println("someone request " + request)
		name := r.URL.Query().Get("name")
		arg := r.URL.Query().Get("arg")
		argsSlice := strings.Split(arg, " ")
		// fmt.Println(argsSlice)
		cmd := exec.Command(name, argsSlice...)
		// fmt.Println(cmd.Path)
		// fmt.Println(cmd.Args)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println("err")
			fmt.Println(err)
		}
		fmt.Println(out.String())
		fmt.Println("command finished.")
		fmt.Fprintf(w, "%s", out.String())
	})
	http.HandleFunc("/ReadDir", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		nexus.ReadDir(w, r)
	})
	http.HandleFunc("/Select", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		nexus.Select(w, r)
	})
	http.HandleFunc("/Select1", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		nexus.Select1(w, r)
	})
	http.HandleFunc("/Select7", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		nexus.Select7(w, r)
	})
	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Println("someone request exit")
		os.Exit(0)
	})
	http.HandleFunc("/web/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		path := r.URL.Query().Get("path")
		http.ServeFile(w, r, path)
	})
	port := ":17000"
	s := &http.Server{
		Addr: "" + port,
	}
	fmt.Println("server at " + s.Addr)
	err := s.ListenAndServe()
	// err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("ListenAndServe: ", err))
	}
}
func SetAccessControlAllowOriginAndPrintRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	request, _ := url.QueryUnescape(r.URL.RequestURI())
	fmt.Println("someone request " + r.Method + " " + request)
}
func cocos(name string) string {
	path := "D:\\cocos2d-x-3.4\\tools\\cocos2d-console\\bin\\cocos"
	// path = "C:\\Users\\G\\Documents\\Visual Studio 2012\\Projects\\Win32Project7\\Debug\\Win32Project7"
	cmd := exec.Command(path, "new", name, "-p", "com.yourcompany."+name, "-l", "cpp", "-d", "D:\\")
	//cmd.Dir = "C:\\"
	fmt.Println(cmd.Path)
	fmt.Println(cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(out.String())
	fmt.Println("create finished.")
	slnPath := "D:\\" + name + "\\proj.win32\\" + name + ".sln"
	exec.Command("explorer.exe", "/select,", slnPath).Run()
	return "done"
}
