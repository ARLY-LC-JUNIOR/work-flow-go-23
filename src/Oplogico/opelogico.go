package main

import "fmt"

func main() {
	x := 18

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("X é multiplo de 2 ou de 3 !!")
	} else {
		fmt.Println("X não é multiplo de 2 e nem 3!!")
		fmt.Println("X não é multiplo de 2 e nem 3!!")
	}
}
