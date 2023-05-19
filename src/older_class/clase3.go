package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c int) int {
	return a + b + c
}

func returnValues(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	value := returnValues(10)
	fmt.Println(value)

	value1, value2 := doubleReturn(2)
	fmt.Println(value1, value2)
}
