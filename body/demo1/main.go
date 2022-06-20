package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/aa", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%#v\n", r.URL)
		fmt.Printf("%#v\n", r.Host)
	})
	http.ListenAndServe(":8081", nil)
}
