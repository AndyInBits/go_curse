package main

import (
	pk "curso_golang_platzi/src/mypkg"
	"fmt"
)

func main() {
	var myCar pk.Car
	myCar.Brand = "Ford"
	myCar.Year = 2020
	fmt.Println(myCar)

	var myOtherCar pk.car
	fmt.Println(myOtherCar)
}
