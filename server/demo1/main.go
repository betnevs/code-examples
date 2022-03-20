package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("172.20.10.5"))
	})
	http.ListenAndServe("172.20.10.5:8080", nil)
	// 172.20.10.5
}
