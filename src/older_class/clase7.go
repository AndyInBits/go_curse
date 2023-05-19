package main

import "fmt"

func main() {

	// Defer
	defer fmt.Println("Hello, 世界")
	fmt.Println("Mundo, 世界")

	// Continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// Break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

}
