package main

import (
	"fmt"
	pk "go/src/curso_golang/src/mypackage"
	//pk "go/src/curso_golang/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)
}

// para utilizar alias en los paquetes vamos a poner el alias
// con el que queremos trabajar del lado izquierdo asi como esta pk
// asi cuando creemos algo de ese paquete solo lo llamaremos por
// pk.atributo o pk.metodo
