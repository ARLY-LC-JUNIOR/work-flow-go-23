// Declarar meu pacote principal utilizando operador curto ou Operador Marmota e Explorando Tipos
package main

//Importar uma função fmt
import "fmt"

//Declaracao da avriavel const do ponto de ebulicao da agua em F
const ebulicaoF = 212.0

//função principal
func main() {

	tempF := ebulicaoF            //Variável do valor da temperatura em grau °F
	tempC := (tempF - 32) * 5 / 9 //conversãod e °F para °C
	//Queremos que apareça na tela o resultado
	fmt.Printf("A temperatura de ebulição da água em °F = %g e temperatura de ebulição da água em °C = %g. ", tempF, tempC)
	//Esperamos que apareça F = 212 2 C = 100
}
