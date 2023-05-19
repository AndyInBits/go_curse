package main

import "fmt"

func main() {
	//Println
	fmt.Println("Hello World")

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("nombre: %T\n", nombre)
	fmt.Printf("cursos: %T\n", cursos)

}
