package main

import (
	"fmt"
	"github.com/marmotedu/errors"
	code "github.com/marmotedu/sample-code"
)

func main() {
	if err := bindUser(); err != nil {
		fmt.Println("----------------%s-----------------\n")
		fmt.Printf("%s\n\n", err)

		fmt.Println("----------------%v-----------------\n")
		fmt.Printf("%v\n\n", err)

		fmt.Println("----------------%-v-----------------\n")
		fmt.Printf("%-v\n\n", err)

		fmt.Println("----------------%+v-----------------\n")
		fmt.Printf("%+v\n\n", err)

		fmt.Println("----------------%#-v-----------------\n")
		fmt.Printf("%#-v\n\n", err)

		fmt.Println("----------------%#+v-----------------\n")
		fmt.Printf("%#+v\n\n", err)

		if errors.IsCode(err, code.ErrEncodingFailed) {
			fmt.Println("this is a ErrEncodingFailed")
		}

		if errors.IsCode(err, code.ErrDatabase) {
			fmt.Println("this is a ErrDatabase")
		}

		fmt.Println(errors.Cause(err))
	}

}

func bindUser() error {
	if err := getUser(); err != nil {
		return errors.WrapC(err, code.ErrEncodingFailed, "encoding user failed")
	}
	return nil
}

func getUser() error {
	if err := queryDatabases(); err != nil {
		return errors.Wrap(err, "get user failed")
	}
	return nil
}

func queryDatabases() error {
	return errors.WithCode(code.ErrDatabase, "user nog found")
}
