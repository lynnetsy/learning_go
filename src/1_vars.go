package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi", pi)
	fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	// 3 tipos de declarar dichas variables
	base := 12 // Se declara y se hace el tipado
	var altura int = 14
	var area int

	fmt.Println("Declaracion Variables:", base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("Zero Values:", a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area Cuadrado:", areaCuadrado)
}
