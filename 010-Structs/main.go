package main

import "fmt"

type User struct {
	name   string
	age    int8
	adress AdressData
} // Serve para criar a tipagem de algo parecido com um objeto, mas não é um objeto para o GO.

type AdressData struct {
	street string
	number int16
	city   string
}

func main() {
	// Declarando a variável e depois atribuindo separadamente seus valores.
	var user01 User
	user01.name = "João"
	user01.age = 30

	fmt.Println(user01)

	// Declarando a variável e atribuindo os valores de uma vez.
	var user02 User = User{name: "Maria", age: 25}

	fmt.Println(user02)

	// Declarando a variável e atribuindo os valores de uma vez.
	user03 := User{name: "Maria", age: 25}

	fmt.Println(user03)

	// Criando usuários com dados faltando, o retorno dos dados que faltam sempre será o valor nulo do tipo da variável
	user04 := User{name: "Maria"}

	fmt.Println(user04) // name: "Maria", age: 0

	user05 := User{age: 10}

	fmt.Println(user05) // name: "", age: 10

	user06 := User{name: "Wesley", age: 21, adress: AdressData{street: "Rua dos Bobos", number: 21, city: "São Paulo"}}

	fmt.Println(user06) // name: "Wesley", age: 21, adress: {street: "Rua dos Bobos", number: 21, city: "São Paulo"}
}
