package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos de números inteiros int8, int16, int32, int64
	// Ou somente int para ele utilize a arquitetura do computador

	var numero int8 = 100
	fmt.Println(numero)

	numero2 := 100
	fmt.Println(numero2)

	// Tipo uint não suporta valores negativos.

	var numero3 uint = 100
	fmt.Println(numero3)

	// alias , alguns tipos possuem um alias , int32 = rune , uint8 = byte
	var numero4 rune = 1020
	var numero5 byte = 102
	fmt.Println(numero4, numero5)

	// Tipos de dados reais float32, float64
	var numeroReal float32 = 123.45
	var numeroReal2 float32 = 12253.45
	fmt.Println(numeroReal, numeroReal2)

	// FIM STRINGS valor padrão da variavel caso ela não seja inicializada
	var texto string
	fmt.Println(texto)

	// Type Boolean
	var booleano1 bool = false
	fmt.Println(booleano1)

	// Type error , o valor 0 dele é nil (nulo)
	var erro error
	fmt.Println(erro)

	// Criar um erro no go, necessário utilização do pacote errors

	var erro_interno error = errors.New("Teste de erro")
	fmt.Println(erro_interno)

}
