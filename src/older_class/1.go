package main

import "fmt"

func main() {
	const pi float64 = 3.14159265
	const pi2 = 3.15

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Suma

	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)
	// Reto

	// Area rectangulo
	baseRectangulo := 20
	alturaRectangulo := 10
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area rectangulo:", areaRectangulo)

	// Area trapecio
	baseMayorTrapecio := 20
	baseMenorTrapecio := 10
	alturaTrapecio := 30
	areaTrapecio := ((baseMayorTrapecio + baseMenorTrapecio) * alturaTrapecio) / 2
	fmt.Println("Area trapecio:", areaTrapecio)

	// Area circulo
	const piCirculo = 3.14159265
	var radioCirculo float64 = 10
	areaCirculo := piCirculo * (radioCirculo * radioCirculo)
	fmt.Println("Area circulo:", areaCirculo)

	// Area poligono
	numeroLadosPoligono := 6
	longitudLadosPoligono := 5
	perimetroPoligono := numeroLadosPoligono * longitudLadosPoligono
	fmt.Println("Perimetro poligono:", perimetroPoligono)

}
