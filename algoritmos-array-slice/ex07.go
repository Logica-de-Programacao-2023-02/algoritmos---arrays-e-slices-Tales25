//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro.
//Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.

package main

import "fmt"

func main() {

	//declaracao slice
	var sliceInt = []int{2, 3, 5, 7, 11}

	//declaracao variavel
	var num int

	//declaracao de variavel booleana (true or false)
	var tester bool

	//coleta do numero informado
	fmt.Print("Escolha um número: ")
	fmt.Scan(&num)

	//
	for _, valor := range sliceInt {

		//se o numero for igual ao valor...
		if num == valor {

			//printa que o slice permanece o mesmo
			fmt.Printf("O slice permanece o mesmo:\n%d", sliceInt)

			//a variavel (booleana) para testar é verdade ou nao recebe true
			tester = true

			//break para sair for/loop caso o numero seja igual ao valor
			break

		}

	}

	//se o tester for o contrario (false), se o numero nao for encontrado
	if !tester {

		//o numero é adicionado ao slice
		sliceInt = append(sliceInt, num)

		//print do "novo" slice (o original modificado)
		fmt.Printf("O novo slice ficou:\n%d", sliceInt)

	}
}
