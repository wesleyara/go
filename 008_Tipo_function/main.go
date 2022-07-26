package main

import "fmt"

func sum(n1 int16, n2 int16) int16 {
	return n1 + n2
} // Criando função de soma

func returnText(text string) string {
	return text
} // Criando função com retorno

func calcMath(n1 int8, n2 int8) (int8, int8) {
	sum := n1 + n2
	mult := n1 * n2

	return sum, mult
} // Criando função que retorna dois valores

func main() {
	var soma int16 = sum(100, 200)
	fmt.Println(soma)

	var function = func(text string) {
		fmt.Println(text)
	} // Atribuindo uma função a uma variável
	function("Testando minha função")

	result := returnText("Testando retorno da minha função")
	fmt.Println(result)

	sum, mult := calcMath(12, 2) // Como a função retorna dois valores, é necessário declarar duas variáveis
	println(sum, mult)

	sum2, _ := calcMath(12, 2) // Quando não quiser usar um dos retornos, invés de criar uma variável, pode-se usar _
	println(sum2)
}
