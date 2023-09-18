//Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas. Em seguida, solicite ao usuário que informe
//um índice de linha e outro de coluna e imprima o valor armazenado nessa posição da matriz.

package main

import "fmt"

func main() {

	arrayMatriz := [3][2]int{{0, 1}, {2, 3}, {4, 5}}

	var i, j int

	fmt.Print("Escolha uma linha [0, 1, 2]: ")
	fmt.Scan(&i)

	fmt.Print("Escolha uma coluna [0, 1]: ")
	fmt.Scan(&j)

	fmt.Printf("O número da linha %d e coluna %d é %d", i, j, arrayMatriz[i][j])

}
