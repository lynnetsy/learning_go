package main

import "fmt"

func main() {
	// For Condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For While
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For Forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
