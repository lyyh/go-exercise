package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wc sync.WaitGroup

func HelloClient(client *http.Client, url string) {
	defer wc.Done()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
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
	url := "http://www.baidu.com"
	wc.Add(10)
	for i := 0; i < 10; i++ {
		go HelloClient(client, url)
	}
	wc.Wait()
}
