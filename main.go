package main

import (
	"chatgpt/api"
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	//http.HandleFunc("/", indexHandler)
	//fs := http.FileServer(http.Dir("static"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/chat/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			// 设置跨域访问权限
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			data := []byte("Hello, world!")
			w.Write(data)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			// 处理 Login 请求
			api.LoginHandler(w, r)
		}
	})
	// 添加路由信息
	http.HandleFunc("/chat/*", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "OPTIONS" {
			// 设置跨域访问权限
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			data := []byte("Hello, world!")
			w.Write(data)
		} else {
			api.ProtectedHandler(w, r)
		}

	})

	fmt.Println("Listening on :8086...")
	http.ListenAndServe(":8086", nil)
}
