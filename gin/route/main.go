package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/aa", func(c *gin.Context) {
		fmt.Println("one step")
	}, func(c *gin.Context) {
		fmt.Println("two step")
	})
	g.GET("/bb", func(context *gin.Context) {

	})
	g.GET("/bac", func(c *gin.Context) {

	})
	g.Run(":8081")
}
