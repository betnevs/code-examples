package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			c.L.Lock()
			ready++
			c.L.Unlock()
			log.Printf("%d ready", i)
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("cp wake")
	}
	c.L.Unlock()
	log.Println("start 1, 2, 3")
}
