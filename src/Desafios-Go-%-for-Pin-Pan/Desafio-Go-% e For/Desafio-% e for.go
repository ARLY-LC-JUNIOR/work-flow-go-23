//Criar um programa na linguagem GO que trabalhe com operador % e for
//Criar um código que exiba todos números entre um a 100 que são divisíveis por 3

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
