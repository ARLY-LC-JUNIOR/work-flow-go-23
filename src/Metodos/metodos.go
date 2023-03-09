//metodo é uma função associada a um tipo particular.
//Isto é em POO Programação Orientada a objetos, objeto é um valor(variavel) e o metodo
//é a função associada a esse objeto.

package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

// este metodo 'área' possui um tipo 'retangulo'
func (r *retangulo) área() int {
	return r.comprimento * r.altura
}

//Metodos podem ser definidos por qualquer tipo de receptor
//ponteiro ou valor aqui temos um exemplo do receptor de valor.
func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func main() {
	r := retangulo{comprimento: 50, altura: 15}
	//aqui será chamado os 2métodos definidos para a nossa estrutura.
	fmt.Println("Área: ", r.área())
	fmt.Println("Perimetro:", r.perimetro())
}
