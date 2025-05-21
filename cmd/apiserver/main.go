package main

import (
	"fmt"
	"quotation_book/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	fmt.Println(config)
	if err := apiserver.Start(config); err != nil {
		panic(err)
	}
}
