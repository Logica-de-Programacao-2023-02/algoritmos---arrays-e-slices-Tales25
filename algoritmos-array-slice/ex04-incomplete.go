//Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos.
//Em seguida, imprima o Slice e a soma dos valores armazenados.

package main

import "fmt"

func main() {

	//declaracao slice
	inteirosSlice := []int{}

	//declaracao de variaveis
	var sliceTamanho, sliceElementos int

	//coleta do tamanho do slice
	fmt.Print("Insira o tamanho desejado par ao slice: ")
	fmt.Scan(&sliceTamanho)

	//o inteirosSlice recebe um slice inteiro (criado agora) com sliceTamanho de tamanho
	inteirosSlice = make([]int, sliceTamanho)

	//
	for i := 0; i < len(inteirosSlice); i++ {

		//
		fmt.Printf("Escolha o %i elemento: ", i)
		fmt.Scan(sliceElementos)

	}

	fmt.Print(inteirosSlice)

}