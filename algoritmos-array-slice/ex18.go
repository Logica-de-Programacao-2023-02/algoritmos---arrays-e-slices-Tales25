FORMA CERTA
//Escreva um programa que leia um número inteiro positivo n e gere um array com os n primeiros números primos.

package main

import "fmt"

func main() {

	//criacao do slice (inicialmente vazio)
	slicePrimes := []int{}

	//declaracao variaveis (quantidade de numeros primos e os proprios numeros primos)
	//como os numeors primos comecam em 2, vou atribuir esse valor
	var qtdPrimes, num int
	num = 2

	//variavel para verificar a veracidade;
	var isPrime bool

	//coleta da quantidade de numeros do slice
	fmt.Print("Escolha a quantidade de números primos: ")
	fmt.Scan(&qtdPrimes)

	//enquanto o tamnho do slice for menor do que qtdPrimes (a quantidade de numeros primos)
	for len(slicePrimes) < qtdPrimes {

		//como o num começa como 2, então isPrime é verdadeiro
		isPrime = true

		//para i==2; enquanto i for menor doq o num; adicionar 1 ao i
		for i := 2; i < num; i++ { /* i==2; num==2; i < num (2 < 2) => como 2 não é menor doq 2
			   só não vai entrar nesse loop */

			//se o resto da divisao do num por i (contador) for 0...
			if num%i == 0 {

				//...entao isPrime recebe falso, pq num nao e primo
				isPrime = false

				//quebra o loop porque ja foi descoberto isPrime e falso
				break

			}
		}

		//se isPrime true?...
		if isPrime {

			//...adicionar o numero ao slice
			slicePrimes = append(slicePrimes, num)

		}

		//adicionar um ao numero e fazer dnv se o tamanho do slice for menor doq a qtd de numeros
		num++

	}

	//print do slice
	fmt.Printf("Os %d primeiros números primos são:\n%d", qtdPrimes, slicePrimes)

}

################################################################################################################
FORMA ERRADA

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
