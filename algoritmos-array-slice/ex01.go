//Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.

package main

import "fmt"

func main() {

	//"declaracao rapida" do array
	arrayInteiros := [3]int{10, 15, 20}

	//declaracao da variavel q armazenarei a soma dos elm do array
	var soma int

	//
	for i := 0; i < len(arrayInteiros); i++ {

		soma += arrayInteiros[i]

	}

	fmt.Print("A soma dos valores armazenados no array é: ", soma)

}
