package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Array tem tamanho pré definido e não pode ser ajustado
	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 100
	fmt.Println(array1)

	array2 := [5]string{"posicao1", "posicao2", "posicao3"}
	fmt.Println(array2)

	// Slice possuí tamanho dinamico
	slice := []int{10, 29, 29, 3, 18}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice = append(slice, 12)
	fmt.Println(slice)

	// Arrays internos

	slice2 := make([]float32, 10, 15)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // length
	fmt.Println(cap(slice2)) // capacidade
	// apesar de definirmos a sua capacidade como 15  o slice não tem limite,
	// caso você faça appends e tente estoura-lo, ele dobra de tamanho automaticamente.

}
