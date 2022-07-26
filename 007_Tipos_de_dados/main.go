package main

import "fmt"

func main() {
	var numero int8 = 100 // o tipo int contém diversos subtipos, como int8, int16, int32 e int64, mudando somente o número de bits suportados.
	fmt.Println(numero)

	var numero2 uint32 = 100 // o tipo uint não permite utilizar sinais negativos, portanto, o valor -100 não é permitido.
	fmt.Println(numero2)

	// alias (apelidos)
	var numero3 rune = 12345 // rune === int32
	fmt.Println(numero3)

	var numero4 byte = 123 // byte === int8
	fmt.Println(numero4)

	var numero5 float32 = 20.3 // floating de 32 bits
	fmt.Println(numero5)

	var numero6 float64 = 201240.3 // floating de 64 bits
	fmt.Println(numero6)

	var string1 string = "String 1" // string
	fmt.Println(string1)

	var boolean1 bool = false // boolean
	fmt.Println(boolean1)

	var voidString string // string vazia
	fmt.Println(voidString)

	var neutralNumber int // número neutro com relação a adição
	fmt.Println(neutralNumber)

	var erro error // erro nulo
	fmt.Println(erro)
}
