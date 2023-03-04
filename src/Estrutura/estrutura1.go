//type(      )struct
//Estruturas = struct são coleções de campos
//serve para agrupar dados
//formar registros

package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(pessoa{"Junior", 50})
	fmt.Println(pessoa{"Cecilia", 45})
	fmt.Println(pessoa{"Junior e Cecilia são casados", 7})
}
