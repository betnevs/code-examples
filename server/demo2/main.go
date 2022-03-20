package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("127.0.0.1"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
	// 172.20.10.5
}
