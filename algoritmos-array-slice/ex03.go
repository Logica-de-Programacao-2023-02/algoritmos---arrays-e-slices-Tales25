//Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.

package main

import "fmt"

func main() {

	//decalracao array
	arrayFloats := [4]float64{1, 2, 3, 4}

	//declaracao de uma variavel para armazenar o prdouto
	var produto float64

	//atribuindo o valor de 1 (pq tudo multiplicado por 0 é 0)
	produto = 1

	//para passar por cada elemento / index de cada elemento
	for i := 0; i < len(arrayFloats); i++ {

		//para produto receber produto vezes o elemento da posicao i
		produto *= arrayFloats[i]

	}

	fmt.Println(arrayFloats)
	fmt.Print("O produto dos elementos do array é ", produto)

}
