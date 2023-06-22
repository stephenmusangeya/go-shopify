package main

import (
	"log"

	"github.com/stephenmusangeya/go-shopify/admin"
)

func main() {
	client := admin.NewClient("store", "token")

	response, err := client.AccessScope(1)
	if err != nil {
		panic(err)
	}

	log.Println(response)
}
