package main

import (
	"example/greetings"
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	message, err := greetings.Hello("Bruce")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
