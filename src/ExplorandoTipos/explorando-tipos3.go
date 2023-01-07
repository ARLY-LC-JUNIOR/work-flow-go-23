// Declarar meu pacote principal utilizando operador curto ou Operador Marmota e Explorando Tipos
package main

//Importar uma função fmt
import "fmt"

//Declaracao da avriavel const do ponto de ebulicao da agua em F
const ebulicaoF = 212.0

//função principal
func main() {

	tempFF := 32
	tempFC := 0
	tempF := ebulicaoF            //Variável do valor da temperatura em grau °F
	tempC := (tempF - 32) * 5 / 9 //Conversão de F para C
	//Queremos que apareça na tela o resultado
	fmt.Printf("A temperatura de ebulição da água em °F = %v (%T), e temperatura de ebulição da água em °C = %v (%T).", tempF, tempF, tempC, tempC)
	fmt.Printf("A temperatura de fusão da água em °F = %v (%T), e temperatura de fusão da água em °C é = %v (%T).", tempFF, tempFF, tempFC, tempFC)
	//Esperamos que apareça F = 212 2 C = 100

}
