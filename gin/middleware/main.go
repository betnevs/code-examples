package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.Use(func(c *gin.Context) {
		fmt.Println("one start")
		c.Next()
		fmt.Println("one end")
	})
	g.GET("/", func(c *gin.Context) {
		fmt.Println("handle")
		c.Abort()
	}, func(c *gin.Context) {
		fmt.Println("two start")
		c.Next()
		fmt.Println("two end")
	})
	g.Run(":8081")
}
