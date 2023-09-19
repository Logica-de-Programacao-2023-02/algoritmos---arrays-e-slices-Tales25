//Escreva um programa que leia um número inteiro positivo n e gere um array com os n primeiros números primos.

package main

import "fmt"

func main() {

	//criacao do slice (inicialmente vazio)
	slicePrimes := []int{}

	//declaracao variavel
	var x, number int

	number = 1

	//variavel booleana para verdadeiro ou falso (atribuindo falso de comeco)
	var primeTester bool
	primeTester = false

	//coleta da variavel
	fmt.Print("Escolha um número: ")
	fmt.Scan(&x)

	//atribuindo a quantidade N de elementos ao slice de numeros primos
	//slicePrimes = make([]int, n)

	//teste para ver como esta o slice
	fmt.Printf("Primeiro teste: %d\n", len(slicePrimes))

	//loop
	for x > len(slicePrimes) {

		//
		if number%2 != 0 && number%3 != 0 || number == 2 || number == 3 {

			//
			primeTester = true

		}

		//
		if number == 1 {

			//
			primeTester = false

		}

		if primeTester == true {

			slicePrimes = append(slicePrimes, number)

		}

		//
		number++

		fmt.Println(len(slicePrimes))

	}

	fmt.Printf("\nO slice com %d números primos ficou:\n%d", x, slicePrimes)

}
