//Faça um programa que leia dois arrays de inteiros de tamanho n e gere um terceiro
//array que seja a soma dos dois primeiros.

package main

import "fmt"

func main() {

	array1 := [5]int{1, 2, 3, 4, 5}
	array2 := [5]int{2, 3, 5, 7, 11}
	slice := []int{}

	for i := 0; i < 5; i++ {

		slice = append(slice, array1[i]+array2[i])

	}

	fmt.Printf("A soma dos arrays %d e %d ficou:\n%d", array1, array2, slice)

}
