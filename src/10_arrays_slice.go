package main

import "fmt"

func main() {
	// Para declarar un array vamos a declarar la variable, el tipo de
	// dato y la dimension, en el 4 le estamos indicando el # de
	// elementos que va a tener este arreglo.
	var array [4]int
	// aca imprimira -> [0 0 0 0]
	fmt.Println(array)
	// Para modificar los elementos de dicho array
	array[0] = 1
	array[1] = 2
	// cap indica la capacidad maxima del array
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append -> agrega un elemento al final del slice
	slice = append(slice, 7)
	fmt.Println(slice)

	// Para recorrer el slice
	for _, valor := range slice {
		fmt.Println(valor)
	}
}
