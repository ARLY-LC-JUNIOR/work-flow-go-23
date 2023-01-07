// Hand's 2 - Estrutura de controles - Se o numeno é PAR ou ÍMPAR//
// O resto que é representado pelo operador%//
package main

import "fmt"

func main() {
	//se i % 2 ==0
	//imprimir que é n° par//
	//se não//
	//imprimir n° ímpar//
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Ímpar")
		}
	}
}
