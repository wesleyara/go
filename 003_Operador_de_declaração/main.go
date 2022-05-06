package main

import "fmt"

var y = "Hello" // Variável de escopo global sempre será declarada utilizando o operador =

func main() {
	x := 42 // Declara a variável no escopo, para novos valores sempre usar :=

	fmt.Printf("x: %v, %T\n", x, x) // Imprime o valor e o tipo da variável
	fmt.Printf("y: %v, %T\n", y, y)

	x = 21 // Atribui ao valor novo para a variável, não é possível atribuir tipos diferentes a variável que já foi declarada

	// O operador = não declara uma nova variável, apenas reatribui o seu valor.

	x, z := 12, "Hello" // Pode-se usar o := para declarar e atribuir, porém é necessário declarar uma nova variável, senão não é possível atribuir valores diferentes a variável que já foi declarada utilizando :=.
	
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
}
