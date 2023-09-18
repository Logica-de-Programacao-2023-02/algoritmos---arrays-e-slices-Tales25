//Crie um Array de inteiros com 7 elementos. Solicite ao usuário que informe um número que será adicionado ao
//primeiro e ao último elemento do Array. Imprima o Array resultante.

package main

import "fmt"

func main() {

	arrayInt := [7]int{1, 2, 3, 4, 5, 6, 7}

	var num1, num2 int

	fmt.Println("Array inicial: ", arrayInt)

	fmt.Print("Escolha um valor para ser adicionado ao primeiro elemento do array: ")
	fmt.Scan(&num1)

	fmt.Print("Escolha um valor para ser adicionado ao ultimo elemento do array: ")
	fmt.Scan(&num2)

	arrayInt[0] += num1
	arrayInt[6] += num2

	fmt.Printf("Array após a soma: %d", arrayInt)

}
