package main

func main() {
	// OPERADORES ARITMÉTICOS; +, -, *, / e %

	soma := 1 + 2          // 3
	subtracao := 2 - 1     // 1
	multiplicacao := 2 * 3 // 6
	divisao := 10 / 5      // 2
	resto := 10 % 3        // 1

	println(soma, subtracao, multiplicacao, divisao, resto, "Operadores aritméticos")

	var num01 int16 = 10
	var num02 int32 = 20
	sum := num01 + int16(num02)
	// Não é possivel somar dois números onde o tipo do dado é diferente, portanto é necessário converter o de int32 para int16.

	println(sum)

	println("------------------------------------")

	// OPERADORES DE ATRIBUIÇÃO; = e :=
	var variavel01 string = "String 01"
	variavel02 := "String 02"

	println(variavel01, variavel02)

	println("------------------------------------")

	// OPERADORES RELACIONAIS
	println(1 > 2, ">")   // false
	println(1 >= 2, ">=") // false
	println(1 == 2, "==") // false
	println(1 < 2, "<")   // true
	println(1 <= 2, "<=") // true
	println(1 != 2, "!=") // true

	println("------------------------------------")

	// OPERADORES LÓGICOS
	// Operador && só retornará true se todas condições forem verdadeiras.
	println(1 < 2 && 3 < 4, "&&") // true

	// Operador || retornará true se pelo menos uma das condições forem verdadeiras.
	println(1 < 2 || 3 > 4, "||") // true
	println(1 > 2 || 3 > 4, "||") // false

	// Operador ! é o de negação e retornará sempre o contrário da condição.
	print(!true, "!") // false

	println("------------------------------------")

	// OPERADORES UNÁRIOS
	num03 := 9
	num03++     // num03 + 1 = 10
	num03 += 15 // num03 + 15 = 25
	num03 -= 15 // num03 - 15 = 10
	num03 *= 2  // num03 * 2 = 20
	num03 /= 2  // num03 / 2 = 10
	num03 %= 4  // num03 % 4 = 2
	println(num03)

	println("------------------------------------")

	// OPERADORES TERNÁRIOS

	var texto string
	if num03 > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	println(texto)
}
