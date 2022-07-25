package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" // Declaração de variável com type explicito
	variavel2 := "Variável 2"           // Declaração de variável com type implicito (inferencia de tipo)

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Outras formas de criar variáveis:
	var (
		variavel3 string = "Variável 3"
		variavel4        = "Variável 4"
	) // Pode criar var sem tipo explicito
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	// Declarando constantes
	const constante1 = "Constante 1"
	println(constante1)

	const constante2 string = "Constante 2"
	println(constante2)

	variavel5, variavel6 = variavel6, variavel5 // Invertendo valores de variáveis
	println(variavel5, variavel6)
}
