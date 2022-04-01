package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"

	"github.com/betNevS/code-examples/sqlc/demo1/tutorial"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("end!!!")
}

func run() error {
	ctx := context.Background()
	db, err := sql.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return fmt.Errorf("cannot connect mysql: %w", err)
	}

	queries := tutorial.New(db)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return fmt.Errorf("cannot query authors: %w", err)
	}
	log.Println(authors)

	result, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "yangjie2",
	})
	if err != nil {
		return err
	}

	insertAuthorID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertAuthorID)

	fetchedAuthor, err := queries.GetAuthor(ctx, insertAuthorID)
	if err != nil {
		return err
	}

	log.Println(reflect.DeepEqual(insertAuthorID, fetchedAuthor.ID))
	return nil
}
