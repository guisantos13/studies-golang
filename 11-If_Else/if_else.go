package main

import "fmt"

func main() {
	fmt.Println("Estruturas de repetição:")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// If init permite inicializar uma variável diretamente no if
	// A variável do if init só existe dentro do escopo do if

	if outroNumero := numero; outroNumero > 15 {
		fmt.Println("Número é maior que 15")
	} else if outroNumero == 15 {
		fmt.Println("Número é 15")
	} else {
		fmt.Println("Número é menor do que 15")
	}

}
