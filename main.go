package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("parse template err:", err)
	}
	msg := "小王子"
	//渲染模版
	t.Execute(w, msg)
}
func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse template err:", err)
	}
	msg := "小王子"
	//渲染模版
	t.Execute(w, msg)

}

func main() {

	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server start failed,err:", err)
		return
	}
}
