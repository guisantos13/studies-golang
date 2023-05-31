package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
	matricula int64
}

func main() {
	pessoa1 := pessoa{"João", "Pedro", 20, 1.80}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Engenharia", "PUCMG", 12300987}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)
}
