package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	// Definindo a estrutura diretamente no loop

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	// Range

	nomes := [3]string{"João", "Davi", "Pedro"}

	for indice, nome := range nomes {

		fmt.Println("Na posição", indice, "O nome é : ", nome)
	}
}
