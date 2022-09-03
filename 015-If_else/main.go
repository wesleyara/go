package main

import "fmt"

func main() {
	number01 := 20

	if number01 > 10 {
		fmt.Println("Esse número é maior que 10")
	} else {
		fmt.Println("Esse número é menor que 10")
	}

	if number02 := 5; number02 > 10 {
		fmt.Println("Esse número é maior que 10")
	} else {
		fmt.Println("Esse número é menor que 10")
	}
}