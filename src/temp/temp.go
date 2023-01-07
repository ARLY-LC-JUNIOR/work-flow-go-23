// declarar meu pacote principal
package main

//importar uma função fmt
import "fmt"

//declaracao da avriavel const do ponto de ebulicao da agua em F
const ebulicaoF = 212.0

//função principal
func main() {

	var tempF = ebulicaoF            //variavel do valor da temperatura em grau F
	var tempC = (tempF - 32) * 5 / 9 //conversãod e F para C
	//queremos que apareça na tela o resultado
	fmt.Println("A temperatura de ebulição da água em °F = ", tempF)
	fmt.Println("A temperatura de ebulição da água em °C = ", tempC)
	//Esperamos que apareça F = 212 2 C = 100
}
