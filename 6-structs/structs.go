package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int16
}

func main() {
	var user1 usuario
	user1.nome = "Guilherme"
	user1.idade = 35
	fmt.Println(user1)

	enderecoHugo := endereco{"Rua cinco", 4}
	user2 := usuario{"Hugo", 3, enderecoHugo}
	fmt.Println(user2)

	user3 := usuario{nome: "Cleber"}
	fmt.Println(user3)
}
