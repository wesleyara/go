package auxiliar

import "fmt"

// Registra uma mensagem no console
func Auxiliar() {
	fmt.Println("Escrevendo do arquivo auxiliar.go")
	escrever2() // Podemos chamar funções que está em arquivos diferentes mas pertencem ao mesmo pacote, mesmo que elas não estejam sendo exportadas (Iniciando com letra maiuscula)
}

// Se uma função começa com letra minuscula ela só pode ser visivel dentro do pacote.