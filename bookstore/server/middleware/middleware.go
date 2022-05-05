package middleware

import (
	"log"
	"mime"
	"net/http"
)

func Logging(nex http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("recv a %s request from %s", req.Method, req.RemoteAddr)
		nex.ServeHTTP(w, req)
	})
}

func Validating(nex http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")
		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if mediaType != "application/json" {
			http.Error(w, "invalid Content-Type", http.StatusBadRequest)
			return
		}
		nex.ServeHTTP(w, req)
	})
}
