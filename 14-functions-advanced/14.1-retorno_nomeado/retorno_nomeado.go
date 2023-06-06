package main

import "fmt"

// Quando declaramos o retorno nomeado na criação da funcão, ele automaticamente retorna as variaveis ao fim
// da execução, apenas declarando o "return".

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
