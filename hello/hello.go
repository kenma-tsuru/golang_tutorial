package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Champ")
	fmt.Println(message)
}