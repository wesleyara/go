package main

import "fmt"

// Os ponteiros são utilizados para pegar o endereço de memória de uma variável, assim evita-se a criação de variáveis duplicadas.

func main() {
	var number01 int = 10
	var number02 int = number01

	fmt.Println(number01, number02)

	number01++

	// Alterar o valor de number01 não altera o valor de number02 e vice versa
	fmt.Println(number01, number02)

	// Ponteiro é um endereço de memória que aponta para um valor
	
	var number03 int = 20
	var number04 *int = &number03 // Utiliza-se o & para obter o endereço de memória dessa variável, ou seja, o endereço de memória da variável number03

	fmt.Println(number03, number04) // 20, o endereço de memória da variável number03

	// Para pegar o valor de um ponteiro, utiliza-se o * e chamamos de desreferenciamento
	fmt.Println(*number04) // 20, ou seja, ele foi no endereço de memória da variável e pegou o valor dela
}
