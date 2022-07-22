package main

import (
	"fmt"

	"github.com/badoux/checkmail" // o comando go get github.com/badoux/checkmail é necessário para instalar o pacote
)

func main() {
	erro := checkmail.ValidateFormat("erro@gmail.com") // checando o formato do email

	fmt.Println(erro)
}

// go mod tidy remove dependencias e modulos desnecessários
