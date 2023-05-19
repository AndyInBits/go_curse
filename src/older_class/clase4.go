package main

import "fmt"

func main() {
	// For condicional
	fmt.Printf("For condicional \n")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Printf("For while \n")

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Printf("For Forever \n")

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 10 {
			break
		}
	}

	fmt.Printf("For reto \n")
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}
