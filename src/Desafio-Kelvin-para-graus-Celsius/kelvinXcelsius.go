// Declarar meu pacote principal
package main

//Importar função fmt
import "fmt"

//Declaração da variável const do ponto de ebulicao da agua em K
const ebuliçãoK = 373.0

//função principal
func main() {

	tempK := ebuliçãoK     //Variavel do valor da temperatura em grau K
	tempC := tempK - 273.0 //Conversão de K para C
	//Queremos que apareça na tela o resultado na mesma linha.
	fmt.Printf("A Temperatura de ebulição da água em K = %g, Temperatura de ebulição da água em °C = %g. ", tempK, tempC)
}
