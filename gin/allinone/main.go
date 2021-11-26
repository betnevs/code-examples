package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Price       int       `json:"price" binding:"gte=0"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"createAt"`
}

type productHandler struct {
	sync.RWMutex
	products map[string]Product
}

func newProductHandler() *productHandler {
	return &productHandler{
		products: make(map[string]Product),
	}
}

func (u *productHandler) Create(c *gin.Context) {
	u.Lock()
	defer u.Unlock()
	var product Product
	// 1、参数解析
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 2、参数校验
	if _, ok := u.products[product.Name]; ok {
		c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("product %s already exist", product.Name)})
		return
	}
	// 3、逻辑处理
	u.products[product.Name] = product
	log.Printf("Register product %s success", product.Name)
	// 4、返回结果
	c.JSON(http.StatusOK, product)
}

func (u *productHandler) Get(c *gin.Context) {
	u.Lock()
	defer u.Unlock()

	product, ok := u.products[c.Param("name")]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("can not found product %s", c.Param("name"))})
	}
	c.JSON(http.StatusOK, product)
}

func router() http.Handler {
	router := gin.Default()
	productHandler := newProductHandler()
	// 路由分组
	v1 := router.Group("/v1", gin.BasicAuth(gin.Accounts{}))
	productv1 := v1.Group("/products")
	productv1.GET(":name", productHandler.Get)
	productv1.POST("", productHandler.Create)
	return router
}

func main() {
	var eg errgroup.Group

	insecureServer := &http.Server{
		Addr:         ":8080",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	secureServer := &http.Server{
		Addr:         ":8443",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	eg.Go(func() error {
		err := insecureServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	eg.Go(func() error {
		err := secureServer.ListenAndServeTLS("server.crt", "server.key")
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
