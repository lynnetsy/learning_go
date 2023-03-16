package main

import "fmt"

func main() {
	if isPair(6) {
		fmt.Println("Es par")
	} else {
		fmt.Println("No es par")
	}

	if isValidUser("lynn", "secret") {
		fmt.Println("Access to your account")
	} else {
		fmt.Println("Your credentials aren't valid")
	}

}

func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(username, pass string) bool {
	return username == "lynn" && pass == "secret"
}
