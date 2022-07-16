package main

import (
	"modulo/auxiliar" // nome do modulo / nome do pacote
	"fmt"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.go")
	auxiliar.Auxiliar() // executando a função exportada do pacote auxiliar
} 


// O comando go build compila o arquivo main.go para um arquivo executável.