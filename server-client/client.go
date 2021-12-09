package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var wc sync.WaitGroup

func HelloClient(client *http.Client, url string, method string) {
	defer wc.Done()
	reqData := "name=ali&age=19"
	req, err := http.NewRequest(method, url, strings.NewReader(reqData))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rep, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(rep.Body)
	rep.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)
}

func main() {
	client := &http.Client{}
	rootUrl := "http://127.0.0.1:8001/login"
	wc.Add(1)
	for i := 0; i < 10; i++ {
		fmt.Print(111)
		go HelloClient(client, rootUrl, "POST")
	}
	wc.Wait()
}
