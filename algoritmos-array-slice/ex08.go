//Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor. Remova todas as
//ocorrências desse valor no Slice e imprima o resultado.

package main

import "fmt"

func main() {

	//declaracao do slice
	sliceStr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	//declaracao variavel
	var letra string

	//coleta da letra
	fmt.Print("Escolha uma letra: ")
	fmt.Scan(&letra)

	//variavel i criada com o valor 0, representando a posicao/index no slice;
	//enquanto ela for menor do que o tamanho do slice, i++ (adicionar 1 ao i)
	for i := 0; i < len(sliceStr); i++ {

		//sliceStr[i] é o valor/string contido na posicao/index i;
		//se esse valor for igual a letra digitada...
		if sliceStr[i] == letra {

			//...o slice vai receber os mesmos valores até a posicao/index i (sem inclui-la),
			//e os valores da posicao/index seguinte (i+1; incluindo-na) até a ultima posicao
			sliceStr = append(sliceStr[:i], sliceStr[i+1:]...)

		}

	}

	//print do slice
	fmt.Printf("O slice ficou:\n%s", sliceStr)

}
