package main

import (
	"fmt"

	"helloworld/greetings"
)

func main() {
	message := greetings.Hello("Bla")
	fmt.Println(message)
}
