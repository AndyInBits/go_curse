package main

import "fmt"

func reverseString(text string) string {
	var reverse string
	for i := len(text) - 1; i >= 0; i-- {
		reverse += string(text[i])
	}
	return reverse
}

func palindrome(text string) bool {
	fmt.Println(text)
	return text == reverseString(text)
}
func main() {

	slice := []string{"holi", "mundo", "cruel"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	palin := "ama"
	result := palindrome(palin)
	fmt.Println(result)

}
