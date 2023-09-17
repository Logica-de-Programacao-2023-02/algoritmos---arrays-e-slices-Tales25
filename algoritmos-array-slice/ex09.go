//Crie um Array de floats com 6 elementos. Solicite ao usuário que informe um número que será adicionado em
//todas as posições do Array. Imprima o Array resultante.

package main

import "fmt"

func main() {

	//craicao do array
	arrayFloats := [6]float64{0.5, 1.5, 2.5, 3.5, 4.5, 5.5}

	//declaracao variavel (tem que ser float tmb, nao pode misturar tipos de variaveis)
	var num float64

	//coleta do numero
	fmt.Print("Escolha um número: ")
	fmt.Scan(&num)

	fmt.Printf("%.2f\n", arrayFloats)

	//loop para somar o numero aos valores armazenados no array
	//i recebe 0 e enquanto o i form menor do que o range/tamanho do array o i recebe + 1;
	//o i é usado como o index dos valores contidos no array (ex: i = 1 => segundo elemento)
	for i := 0; i < len(arrayFloats); i++ {

		//somar o numero ao valor contido na posicao/index i
		arrayFloats[i] += num

	}

	//print do array após a soma
	fmt.Printf("Somando seu número a cada elemento do array, ele ficou:\n%.2f", arrayFloats)

}
