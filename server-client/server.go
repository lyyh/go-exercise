package main

import (
	"fmt"
	"net/http"
	"time"
)

func login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(string("post")))
	req.ParseForm()
	fmt.Println("login")
	fmt.Println(req.Form["name"])
	fmt.Println(req.Form["age"])
}

func main() {
	server := &http.Server{
		Addr:         "127.0.0.1:8001",
		ReadTimeout:  2 * time.Second, // 读超时
		WriteTimeout: 2 * time.Second, // 写超时
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)

	server.Handler = mux
	server.ListenAndServe()
}
