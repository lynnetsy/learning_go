package main

import "fmt"

func main() {
	// Declaracion de variables
	helloMsg := "Hello"
	worldMsg := "World"

	//Println (imprime en consola + salto de linea)
	fmt.Println(helloMsg, worldMsg)

	// Printf (Imprime y agrega una funcion extra
	// como variable de entrada)
	// %s = significa que el valor sera un string
	// %d = significa que el valor sera un entero int
	// %v = significa que imprimira el valor, si no tenemos
	// la certeza del tipo de dato que es.
	nombre := "chocolate kinder"
	precio := 20
	fmt.Printf("el %s cuesta %d pesos\n", nombre, precio)

	//Sprintf
	message := fmt.Sprintf("el %s cuesta %d pesos", nombre, precio)
	fmt.Println(message)

	// %T sirve para imprimir el tipo de dato que es una variable
	// Este %T con println no sirve, tambien necesita \n al final.
	fmt.Printf("nombre: %T\n", precio)

}
