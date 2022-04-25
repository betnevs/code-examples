package main

type Handler interface {
	ServeHTTP(c *Context)
	Router
}

type Router interface {
	Route(method string, pattern string, handlerFunc func(ctx *Context))
}
