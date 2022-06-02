package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	setupLogger()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("https://www.baidu.com")
}

func setupLogger() {
	logFileLocation, _ := os.OpenFile("/tmp/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	log.SetOutput(logFileLocation)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Println("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}
