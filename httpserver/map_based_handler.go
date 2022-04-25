package main

import (
	"net/http"
)

var _ Handler = (*HandlerBasedOnMap)(nil)

type HandlerBasedOnMap struct {
	handlers map[string]func(c *Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	key := h.key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found Route"))
	}
}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(c *Context)),
	}
}
