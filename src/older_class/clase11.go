package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	carro := car{brand: "Ford", year: 2020}
	fmt.Println(carro)

	// Otro tipo de declaracion
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2020
	fmt.Println(otherCar)
}
