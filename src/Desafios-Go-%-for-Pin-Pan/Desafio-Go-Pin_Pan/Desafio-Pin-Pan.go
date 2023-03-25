// Criar um código que imprima todos números de um a 100 e nos casos
// citados não devem aparecer os números mas sim o que o participante
// deve falar que são múltiplos de 3 mas no lugar do múltiplo aparecer a palavra PIN
// e nos múltiplos de 5 aparecer no lugar a palavra PAN.
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("PIN ")
		} else if i%5 == 0 {
			fmt.Print("PAN ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}
