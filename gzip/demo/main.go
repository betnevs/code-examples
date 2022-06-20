package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "adfajsidofjaosidjfioasdjniofjasdiojasiodjfioasdjiovjnaiosdnvioasdjnvioandvionaso")
	})

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
