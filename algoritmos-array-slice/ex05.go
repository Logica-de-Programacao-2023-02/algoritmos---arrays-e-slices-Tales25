//Crie um Array bidimensional (Matriz) de inteiros com 3 linhas e 2 colunas. Solicite ao usuário que informe
//os valores de cada elemento da matriz. Em seguida, imprima a matriz resultante.

package main

import "fmt"

func main() {

	//declaracao da matriz
	var matrizInteiros [3][2]int

	//declaracao variaveis
	var matrizValores int

	//loop para as linhas
	for i := 0; i < 3; i++ {

		//loop para as colunas (dentro do loop de linhas)
		for j := 0; j < 2; j++ {

			//coleta de valores
			fmt.Printf("Escolha o valor da posição [%d][%d]: ", i, j)
			fmt.Scan(&matrizValores)

			//atribuicao dos valores a matriz
			matrizInteiros[i][j] = matrizValores

		}

	}

	//print da matriz
	//fmt.Printf("\nSua matriz ficou:\n%d", matrizInteiros)
	for _, ints := range matrizInteiros {

		fmt.Println(ints)

	}

}
