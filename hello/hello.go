package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	log.setPrefix("greetings: ")
	log.setFlags(0)
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)

}