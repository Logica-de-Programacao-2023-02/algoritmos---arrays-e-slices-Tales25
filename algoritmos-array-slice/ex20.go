//Crie um programa que leia um array/slice de inteiros e verifique se ele está ordenado em ordem crescente.

package main

import "fmt"

func main() {

	slice := []int{}

	var qtdElementos, elemento int
	var isGrowing bool //para verificar se esta crescendo(true) ou nao (false)

	fmt.Print("Quantidade de elementos: ")
	fmt.Scan(&qtdElementos)

	for i := 0; i < qtdElementos; i++ {

		fmt.Printf("Elemento %d: ", i)
		fmt.Scan(&elemento)

		slice = append(slice, elemento)

	}

	//para i==0; enquanto i for menor do que o tamanho so slice menos 1; adiconar 1 ao i
	for i := 0; i < len(slice)-1; i++ { //len(slice)-1 para o slice[i+1] não comparar com um elemento inexistente

		isGrowing = true //é crescente (isGrowing=true) até que o if prove o contrário

		//se o elemento sucessor for menor do que o antecessor...
		if slice[i+1] < slice[i] {

			isGrowing = false //...não é crescente (isGrowing=false)
			break             //o loop pode ser quebrado, pq já não é crescente (isGrowing=false)

		}

	}

	if isGrowing { //se for verdadeiro; é como if isGrowing == true

		fmt.Print("Está em ordem crescente")

	} else {

		fmt.Print("Não está em ordem crescente")

	}

}
