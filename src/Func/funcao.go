//Função também é chamada de procedimento ou sub-rotina
//parte código
//ela pega um dado de entrada e dá um dado de saida

package main

import "fmt"

func main() { //irei criar um programa para calcular a média de uma sala
	lista := []float64{98, 93, 77, 82, 83} //lista de notas

	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	fmt.Println(total / float64(len(lista))) // imprmir a media da sala

}
