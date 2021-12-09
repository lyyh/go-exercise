package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var wc sync.WaitGroup

func sendData(client *http.Client, url string, method string, filePath string) {
	defer wc.Done()
	if client == nil {
		log.Fatal("client is nil")
	}
	boundary := "ASSDFWDFBFWEFWWDF"
	if _, err := os.Lstat(filePath); err == nil {
		file, _ := os.Open(filePath)
		defer file.Close()
		data, _ := ioutil.ReadAll(file)
		picData := "--" + boundary + "\n"
		picData = picData + "Content-Disposition: form-data; name=\"userfile\"; filename=" + filePath + "\n"
		picData = picData + "Content-Type: application/octet-stream\n\n"
		picData = picData + string(data) + "\n"
		picData = picData + "--" + boundary + "\n"
		picData = picData + "Content-Disposition: form-data; name=\"text\";filename=\"1.txt\"\n\n"
		picData = picData + string("data=ali") + "\n"
		picData = picData + "--" + boundary + "--"

		req, err := http.NewRequest(method, url, strings.NewReader(picData))
		req.Header.Set("Content-Type", "multipart/form-data;boundary="+boundary)
		if err != nil {
			if rep, err := client.Do(req); err == nil {
				content, _ := ioutil.ReadAll(rep.Body)
				log.Panicln("get response:" + string(content))
				rep.Body.Close()
			}
		}
	} else {
		log.Fatal("file not exist")
	}
}

func main() {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	filePath := "profile.png"
	method := "POST"
	url := "http://127.0.0.1:8000/postdata"
	wc.Add(1)
	go sendData(client, url, method, filePath)
	wc.Wait()
}
