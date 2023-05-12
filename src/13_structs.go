package main

import "fmt"

// Struct es una clase

type car struct {
	brand string
	year  int
}

func main() {
	mycar := car{brand: "Ford", year: 2020}

	fmt.Println(mycar)
	fmt.Println(mycar.brand)
	fmt.Println(mycar.year)

	// Otra forma de instanciar la clase sera

	var otherCar car
	otherCar.brand = "Ferrari"

	fmt.Println(otherCar)
}
