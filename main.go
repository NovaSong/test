package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for {
		fuckthenet()
		time.Sleep(10 * time.Second)
	}
}

func fuckthenet() {
	url := "http://www.google.com"

	for i := 0; i < 10; i++ {

		go DoARequest(url)
		// time.Sleep(1 * time.Second)
	}
}

func DoARequest(url string) {
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	//设置日志输出到 f

	log.SetOutput(f)

	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		log.Println(err)
	}
	if response != nil {
		status := response.StatusCode

		fmt.Println(status)
		log.Println(status)
	}

}
