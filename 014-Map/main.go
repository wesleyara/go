package main

import "fmt"

func main() {
	usuario01 := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	} // Os tipos dentro de um map devem ser iguais, ou seja, todos os valores devem ser do mesmo tipo e todos as chaves devem ser do mesmo tipo

	fmt.Println(usuario01)

	usuario02 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"último":   "Silva",
		},
		"curso": {
			"nome":   "Engenharia de Software",
			"campus": "Campus Rio Vermelho",
		},
	} // Aninhando Maps

	fmt.Println(usuario02)
	delete(usuario02, "nome") // Deletando um item do map
	fmt.Println(usuario02)

	usuario02["profissão"] = map[string]string{
		"nome": "Programador",
	} // Adicionando um item ao map
	fmt.Println(usuario02)
}
