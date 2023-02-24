package main

func main() {
	// laço for simples
	for i := 0; i <= 10; i++ {
		println(i)
	}

	println("done")

	// laço for com break
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		println(i)
	}

	println("done")

	// laço for range
	names := [5]string{"Wesley", "Kleber", "João", "Felipe", "Lucas"}

	for index, value := range names {
		println(index, value)
	} 

	// laço for range com _
	for _, value := range names {
		println(value)
	}

	// laço for range com string
	for index, value := range "Wesley" {
		// é necessário usar o método string para converter o valor para string
		println(index, string(value))
	}

	// laço for no array
	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	// laço for no map
	ages := map[string]int{
		"Wesley": 20,
		"Kleber": 21,
		"João": 22,
		"Felipe": 23,
		"Lucas": 24,
	}

	for key, value := range ages {
		println(key, value)
	}
}