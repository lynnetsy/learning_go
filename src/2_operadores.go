package main

import "fmt"

func main() {
	// Area Cuadrado
	const baseCuadrado int8 = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("area cuadrado", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma", result)

	// Resta
	result = y - x
	fmt.Println("Resta", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion")

	// Division
	result = y / x
	fmt.Println("Division", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo", result)

	// Incremental
	x++
	fmt.Println("Incremental", x)

	// Decremental
	x--
	fmt.Println("Decremental", x)
}
