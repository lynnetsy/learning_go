package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 5, "Hola")

	// Aca asignamos a una variable que ejecute la funcion returnValue
	// asi despues lo imprimimos con la libreria fmt
	value := returnValue(2)
	fmt.Println("Value:", value)

	// Utilizamos una funcion que retorna 2 valores :) por lo que
	// declaramos los 2 valores.
	value1, value2 := doubleReturn(2)
	fmt.Println("value 1 y value 2", value1, value2)

	// Si tenemos una funcion que retorna 2 valores y solo queremos saber
	// uno de sus valores, hacemos lo siguiente declaramos la variable y
	// luego  _ (piso), eso indica a go que solo nos interesa

	value3, _ := doubleReturn(4)
	fmt.Println("Solamente retornamos value3", value3)

}
