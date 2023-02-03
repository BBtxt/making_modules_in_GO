package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	//Get a greeting message and print it out
	message := greetings.Hello("Gladys")
	fmt.Println(message)

}
