package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// with and
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad")
	}

	// with or

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Cumplio una condicion del or")
	}
}
