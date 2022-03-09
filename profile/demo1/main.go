package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			log.Println(Add("http://www.baidu.com"))
			time.Sleep(time.Second)
		}
	}()
	log.Fatal(http.ListenAndServe(":8888", nil))
}
