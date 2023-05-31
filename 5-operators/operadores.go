package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	mult := 10 * 5
	restDiv := 10 % 2

	fmt.Println(soma, sub, div, mult, restDiv)

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	// RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// UNARIOS
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)
}
