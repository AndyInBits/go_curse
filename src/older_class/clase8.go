package main

import "fmt"

func main() {

	var array [4]int
	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos de slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
	fmt.Println(slice[:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
