package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		println("Es 1")
	} else {
		println("No es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 2 {
		println("Es 1 y 2")
	} else {
		println("No es 1 y 2")
	}

	// With or
	if valor1 == 1 || valor2 == 2 {
		println("Es 1 o 2")
	} else {
		println("No es 1 o 2")
	}

	// With switch
	switch valor1 {
	case 1:
		println("Es 1")
	case 2:
		println("Es 2")
	default:
		println("No es 1 ni 2")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("45")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("El numero es %d\n", value)
}
