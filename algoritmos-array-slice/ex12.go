//Crie um Array de inteiros com 5 elementos. Em seguida, crie um novo Slice que contenha apenas os elementos
//do Array que são múltiplos de 3. Imprima o novo Slice.

package main

import "fmt"

func main() {

	arrayInt := [10]int{66, 16, 12, 59, 34, 69, 47, 61, 23, 18}
	sliceInt := []int{}

	for i := 0; i < len(arrayInt); i++ {

		if arrayInt[i]%3 == 0 {

			sliceInt = append(sliceInt, arrayInt[i])

		}

	}

	fmt.Printf("Array: %d\nSlice com os múltiplos de 3: %d", arrayInt, sliceInt)

}
