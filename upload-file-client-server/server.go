package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//DownloadFile download file from client to local.
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	log.Print("aaaa")
	switch r.Method {
	case "GET":
		fmt.Println("GET")
		w.Write([]byte(string("hi, get successful")))
	case "POST":
		fmt.Println("POST")
		r.ParseForm()                             //解析表单
		imgFile, _, err := r.FormFile("userfile") //获取文件内容
		if err != nil {
			log.Fatal(err)
		}
		defer imgFile.Close()

		imgName := "1.png"
		files := r.MultipartForm.File //获取表单中的信息
		for k, v := range files {
			for _, vv := range v {
				fmt.Println(k + ":" + vv.Filename) //获取文件名
				if strings.Index(vv.Filename, ".png") > 0 {
					imgName = vv.Filename
				}
			}
		}

		saveFile, _ := os.Create(imgName)
		defer saveFile.Close()
		io.Copy(saveFile, imgFile) //保存

		w.Write([]byte("successfully saved"))
	default:
		fmt.Println("default")
	}
}

func main() {
	// 创建多路复用处理函数，因为 ServeMux 实现的功能就是将收到的 HTTP 请求的 URL 与注册的路由相匹配，选择匹配度最高的路由的处理函数来处理该请求。
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/postdata", DownloadFile)
	// 创建自定义服务器
	server := &http.Server{
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		Handler:      serveMux,
	}
	err := server.ListenAndServe()
	fmt.Print(err)
	if err != nil {
		fmt.Println("listening port is 8000")
	} else {
		log.Fatal(err)
	}
}
