package main

import "fmt"

// Recebe uma quantidade variavel de parametros.
func soma(numeros ...int) int {

	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}

func main() {
	totalSoma := soma(1, 3, 4, 6, 10, 22)
	fmt.Println(totalSoma)

}
