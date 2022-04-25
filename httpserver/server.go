package main

import "net/http"

type Server interface {
	Start(address string) error
	Router
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnMap()

	var root Filter = func(c *Context) {
		handler.ServeHTTP(c)
	}

	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{Name: name, handler: handler, root: root}
}

func main() {
	server := NewHttpServer("test-server", MetricsFilterBuilder)
	server.Route(http.MethodGet, "/hello", func(ctx *Context) {
		ctx.OkJson("OK")
	})
	server.Start(":8080")
}
