//Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.

package main

import "fmt"

func main() {

	//criacao do slice
	sliceInt := []int{66, 16, 12, 59, 34, 69, 47, 61, 23, 18}

	//declaracao variavel maior e menor
	var maiorNum, menorNum int

	//declaracao do index aqui para ja terem um valor base/comparativo
	i := 0
	maiorNum = sliceInt[i]
	menorNum = sliceInt[i]

	//loop para passar por todos os elementos do array
	for i = 0; i < len(sliceInt); i++ {

		if sliceInt[i] > maiorNum {

			maiorNum = sliceInt[i]

		}

		if sliceInt[i] < menorNum {

			menorNum = sliceInt[i]

		}

	}

	fmt.Printf("O menor número foi %d e o maior %d", menorNum, maiorNum)

}
