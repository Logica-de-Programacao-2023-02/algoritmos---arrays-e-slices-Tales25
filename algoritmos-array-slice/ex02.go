//Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.

package main

import "fmt"

func main() {

	//declaracao slice
	sliceInteiros := []int{1, 2, 3, 4, 5}

	//slice antes
	fmt.Println(sliceInteiros)

	//remocao de um elemento
	sliceInteiros = append(sliceInteiros[:2], sliceInteiros[3:]...)

	//slice depois
	fmt.Print(sliceInteiros)

}
