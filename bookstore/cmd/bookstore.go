package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/betNevS/code-examples/bookstore/server"

	_ "github.com/betNevS/code-examples/bookstore/internal/store"
	"github.com/betNevS/code-examples/bookstore/store/factory"
)

func main() {
	s, err := factory.New("mem")
	if err != nil {
		panic(err)
	}

	bookStoreServer := server.NewBookStoreServer(":8080", s)
	errChan, err := bookStoreServer.ListenAndServe()
	if err != nil {
		log.Println("start failed: ", err)
		return
	}

	log.Println("begin server")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		log.Println("web server run: ", err)
	case <-c:
		log.Println("server is exiting...")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		err = bookStoreServer.Shutdown(ctx)
	}

	if err != nil {
		log.Println("exit error: ", err)
		return
	}

	log.Println("exit ok")
}
