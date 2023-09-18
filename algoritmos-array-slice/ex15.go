//Crie um Array de floats com 10 elementos. Crie um novo Slice que contenha apenas os elementos do Array
//que são maiores que 5. Imprima o novo Slice.

package main

import "fmt"

func main() {

	//criacao do array e do slice
	arrayFloats := [10]float64{2, 3, 5, 7, 11, 17, 2.5, 3.5, 4.5, 5.5}
	sliceFloats := []float64{}

	//print do array inicial
	fmt.Println(arrayFloats)

	//loop para passar por todo o array
	for i := 0; i < len(arrayFloats); i++ {

		//se o valor do index i no array for maior do que 5...
		if arrayFloats[i] > 5 {

			//...o slice (vazio inicialmente) recebe o proprio slice e o indice i do array
			sliceFloats = append(sliceFloats, arrayFloats[i])

		}

	}

	//print do novo slice
	fmt.Printf("O novo slice com valores maiores que 5 ficou:\n%.2f", sliceFloats)

}
