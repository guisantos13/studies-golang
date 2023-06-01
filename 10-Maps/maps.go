package main

import "fmt"

func main() {
	// Não conseguimos utilizar o "." para acessar as chaves como fazemos no struct
	// As chaves e valores devem ser do mesmo tipo

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	// Aninhando MAPS

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia de dados",
			"campus": "UFMG",
		},
	}
	// Deletando chaves do MAP

	delete(usuario2, "nome")
	fmt.Println(usuario2)

}
