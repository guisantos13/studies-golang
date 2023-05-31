package main

import "fmt"

// Declaração de variáveis e identificação dos tipos.

func main() {
	// Declaração com tipos explícitos
	var variavel1 string = "Variável 1"
	//Declaração com inferência de tipo :=
	variavel2 := "Variável 2"

	// Declaração de variáveis em conjunto com tipo explícito

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	// Declaração de variáveis em conjunto com inferência de tipo

	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)

}
