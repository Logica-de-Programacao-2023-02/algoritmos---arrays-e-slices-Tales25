//Crie um Array de inteiros com 10 elementos. Calcule e imprima a soma dos elementos nas posições pares do Array.

package main

import "fmt"

func main() {

	//criacao do array
	arrayInt := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//declaracao da variavel usada para acumular a soma dos valores
	var soma int

	//print do array
	fmt.Println(arrayInt)

	//loop para passar por todos os valores do array;
	//e atribui-los a variavel elementos
	for _, elementos := range arrayInt {

		//soma (inicialmente 0) recebe soma + elementos (do primeiro ate o ultimo elementos)
		soma += elementos

	}

	fmt.Printf("A soma dos valores do array é %d", soma)

}
