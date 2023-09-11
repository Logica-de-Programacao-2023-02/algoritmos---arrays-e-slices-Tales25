//Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos.
//Em seguida, imprima o Slice e a soma dos valores armazenados.

package main

import "fmt"

func main() {

	//declaracao slice
	inteirosSlice := []int{}

	//declaracao de variaveis
	var sliceTamanho, sliceElemento, soma int

	//coleta do tamanho do slice
	fmt.Print("Insira o tamanho desejado do slice: ")
	fmt.Scan(&sliceTamanho)

	//o inteirosSlice recebe um slice inteiro (criado agora) com sliceTamanho de tamanho
	inteirosSlice = make([]int, sliceTamanho)

	//loop (i recebe 0; enquanto i for menor doq o tamanho do slice(i < numero digitado), adicione um ao i)
	for i := 0; i < len(inteirosSlice); i++ {

		//coleta dos elementos do slice
		fmt.Printf("Escolha o %d elemento: ", i)
		fmt.Scan(&sliceElemento)

		//atribuindo o elemento a determinada posicao
		inteirosSlice[i] = sliceElemento

		//somatorio dos elementos
		soma += sliceElemento

	}

	//print do slice
	fmt.Printf("Seu slice ficou:\n%d", inteirosSlice)

	//print da soma de todos os elementos
	fmt.Printf("\nA soma de todos seus elementos é: %d", soma)

}
