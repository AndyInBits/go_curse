package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go func(text string) {
		c <- text
	}("Bye")

	fmt.Println(<-c)
}
