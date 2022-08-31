package main

import "fmt"

type Person struct {
	name     string
	lastname string
	age      int8
	height   float32
}

// Todo Stundets é uma pessoa, logo ele herda as propriedades de Person, portanto não preciso repetir os campos, basta usar o nome do struc que ele herda as propriedades e adicionar novas. Muito semelhante a um "interface Students extends Person"
type Student struct {
	Person
	school string
	curse  string
}

// Com isso posso acessar Students.name e Students.age e etc...

func main() {
	var person01 Person = Person{"João", "Silva", 20, 1.80}

	fmt.Println(person01)

	var student01 Student = Student{Person{"João", "Silva", 20, 1.80}, "IFPI", "Engenharia de Software"}

	fmt.Println(student01.name)
}
