package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	text_lower := strings.ToLower(text)

	for i := len(text_lower) - 1; i >= 0; i-- {
		textReverse += string(text_lower[i])
	}

	if text_lower == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	isPalindromo("Ama")
}
