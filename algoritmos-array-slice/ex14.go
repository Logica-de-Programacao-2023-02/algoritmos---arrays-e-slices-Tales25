//Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de elementos que
//deverão ser trocados de posição. Imprima o Slice resultante.

package main

import "fmt"

func main() {

	//criacao do slice
	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//declaracao variaveis
	var x, y int

	//print do slice
	fmt.Println(sliceInt)

	//coleta dos indices
	fmt.Println("Escolha as posições (de 0 a 7) a serem trocadas:")
	fmt.Scan(&x, &y)

	//print dos valores contidos nos indices
	fmt.Printf("O valor do Primeiro Index é %d e do Segundo Index é %d\n", sliceInt[x], sliceInt[y])

	//variaveis usadas para trocar os valores
	troca1 := sliceInt[x]
	troca2 := sliceInt[y]

	//trocando de fato os valores
	sliceInt[x] = troca2
	sliceInt[y] = troca1
	
	//jeito do professor (da para trocar mais de uma variavel na mesma linha par não dar erro)
	//sliceInt[x], sliceInt[y] = sliceInt[y], sliceInt[x]

	//print do novo slice
	fmt.Print(sliceInt)

}
