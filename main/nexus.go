﻿package main

import (
	"../../nexus"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// nexus.ConnectRedis()
	defer nexus.CloseRedisPool()

	go server()
	var input string
	fmt.Scanln(&input)
	fmt.Println("server closed")
}
func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
		//SetAccessControlAllowOriginAndPrintRequest(w, r)
	})
	//http.Handle("/", http.FileServer(http.Dir("./index.html")))
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
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		fmt.Fprintf(w, "%s", "hi")
	})
	http.HandleFunc("/hi/hi", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		fmt.Fprintf(w, "%s", "hi/hi")
	})
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		nexus.List(w, r)
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
	// http.HandleFunc("/Select", func(w http.ResponseWriter, r *http.Request) {
	// 	SetAccessControlAllowOriginAndPrintRequest(w, r)
	// 	nexus.Select(w, r)
	// })
	http.HandleFunc("/Redis", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		nexus.Redis(w, r)
	})
	http.HandleFunc("/Postgres/Select", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		nexus.PostgresSelect(w, r)
	})
	http.HandleFunc("/MySQL/Select", func(w http.ResponseWriter, r *http.Request) {
		SetAccessControlAllowOriginAndPrintRequest(w, r)
		nexus.MySQLSelect(w, r)
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
	t := time.Now()
	fmt.Println("[" + t.Format("15:04:05") + "] someone request " + r.Method + " " + request)
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
