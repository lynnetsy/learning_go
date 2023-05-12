package main

import "fmt"

func main() {
	// Map sirve para crear distintas vars

	m := make(map[string]int)

	m["pepito"] = 10
	m["Jose"] = 20

	// La forma de visualizar el diccionario no se ve separado por
	// coma si no por un espacio
	fmt.Println(m)

	// Recorrer map
	// sin embargo cuando creamos un arreglo con map
	// se generan de forma concurrente por lo que no hay un orden
	// especifico.
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor, se utilizan las ""
	value := m["Jose"]
	fmt.Println(value)

}
