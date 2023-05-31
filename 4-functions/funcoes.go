package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Função com mais de 1 retorno, caso queira desconsiderar um dos valores ao chamar a funcao utilize o _
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Para desconsiderar um dos valores retornados pela função podemos utilizar o _
	resultado, _ := calculosMatematicos(10, 15)
	fmt.Println(resultado)
}
