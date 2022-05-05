package main

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  handlerFunc
}

func (n *node) ServeHTTP(c *Context) {
	//TODO implement me
	panic("implement me")
}

func (n *node) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	//TODO implement me
	panic("implement me")
}
