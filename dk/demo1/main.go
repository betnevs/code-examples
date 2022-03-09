package main

import (
	"log"
	"net/http"
)

const (
	addr = ":8080"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request path: ", r.URL.Path, ", client ip: ", r.RemoteAddr)
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(addr, nil)
}
