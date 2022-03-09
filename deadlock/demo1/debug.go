package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type SafeMap struct {
	lock sync.Mutex
	data map[string]interface{}
}

var Sm = &SafeMap{data: map[string]interface{}{}}

func main() {
	go worker1()
	go worker2()

	log.Println(http.ListenAndServe("10.16.30.54:8080", nil))
}

func worker1() {
	for {
		fmt.Println("worker1 start lock")
		Sm.lock.Lock()
		defer Sm.lock.Unlock()
		Sm.data["test"] = 1
		fmt.Println("worker1 end write")
		time.Sleep(10 * time.Second)
	}
}

func worker2() {
	for {
		fmt.Println("worker2 start lock")
		Sm.lock.Lock()
		defer Sm.lock.Unlock()
		Sm.data["test"] = 2
		fmt.Println("worker2 end write")
		time.Sleep(10 * time.Second)
	}
}
