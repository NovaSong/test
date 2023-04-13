package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	url := "http://www.google.com"
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	for {
		for i := 0; i < 30000; i++ {
			go func() {
				//处理返回结果
				response, err := client.Do(reqest)
				if err != nil {
					log.Println(err)
				}
				if response != nil {
					status := response.StatusCode
					if status == 200 {
						f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
						if err != nil {
							log.Println(err)
						}
						defer f.Close()
						//设置日志输出到 f
						log.SetOutput(f)
						log.Println(status)
					}
					log.Println(status)
				}
			}()
			go func() {
				for i := 1; i < 10000000; i++ {

				}
			}()
			// time.Sleep(1 * time.Second)
		}
		time.Sleep(10 * time.Second)
		runtime.GC()
	}
}
