//Crie um Array de inteiros com 10 elementos. Crie um novo Slice que contenha apenas
//os elementos pares do Array original. Imprima o novo Slice.

package main

import "fmt"

func main() {

	//criacao do array e do slice
	arrayInt := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceEvensInt := []int{}

	//print do array inicial
	fmt.Println(arrayInt)

	//loop para passar por todos os elementos do array
	for i := 0; i < len(arrayInt); i++ {

		//se o resto da divisao do elemento da posicao i for 0...
		if arrayInt[i]%2 == 0 {

			//...adiciona o elemento ao slice (inicialmente 0)
			sliceEvensInt = append(sliceEvensInt, arrayInt[i])

		}

	}

	//print do slice contendo os numeros pares do array
	fmt.Printf("O novo slice contendo números pares do array ficou:\n%d", sliceEvensInt)

}