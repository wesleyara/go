package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Arrays

	var array01 [5]int // Em Go é necessário sempre especificar o tamanho do array e ele só pode receber dados de tipos iguais, por exemplo, o array ["1", 0, 2] não é possível ser formado.
	array01[3] = 15    // Atribuindo valor ao elemento 3 do array

	fmt.Println(array01)

	array02 := [3]string{"Wesley", "Filipe", "João"} // Criando array por inferencia de tipos, já atribuindo os valores dos campos

	fmt.Println(array02)

	array03 := [...]int{1, 2, 3, 4, 5} // o [...] faz com que ele fixe o tamanho do array de acordo com a quantidade de elementos que ele recebe

	fmt.Println(array03)

	// Slices, os slices são referencias de arrays (como um ponteiro)

	slice01 := []int{1, 2, 3, 4, 5, 6} // Slice é um array que não tem tamanho fixo, ele pode crescer ou diminuir de acordo com a necessidade

	fmt.Println(slice01)

	fmt.Println(reflect.TypeOf(slice01)) // O reflect.TypeOf() retorna o tipo do dado que foi passado como parâmetro
	fmt.Println(reflect.TypeOf(array01)) // Como o slice e o array não são do mesmo tipo, não é possível fazer operações entre eles

	slice01 = append(slice01, 7) // O append() adiciona um elemento ao slice

	fmt.Println(slice01)

	slice02 := array02[0:2] // O slice02 recebe os elementos do array02 que estão no índice 0 e abaixo do indice 2, ou seja, o elemento 1

	fmt.Println(slice02)

	///////////////////////////////////////////////////////////////////////////////////////////////////////

	// Arrays Internos
	slice03 := make([]float32, 10, 11) // O make() cria um array de 15 posições e retorna um slice que retorna as 10 primeiras posições, ou seja, invés de referenciar um array genérico, ele criou e alocou na memória um array e pegou uma referência dele. Caso não passe a capacidade do slice, ele irá pegar o tamanho do slice como capacidade.

	fmt.Println(slice03)
	fmt.Println(len(slice03)) // O len() retorna o tamanho do slice
	fmt.Println(cap(slice03)) // O cap() retorna a capacidade do slice

	slice03 = append(slice03, 2)
	slice03 = append(slice03, 2)

	fmt.Println(slice03)
	fmt.Println(len(slice03))
	fmt.Println(cap(slice03))
	// Ao adicionar itens e passar da capacidade do slice, ele cria um novo array com o dobro da capacidade e copia os dados do array antigo para o novo, ou seja, o slice é um array que cresce dinamicamente
}
