// declarar meu pacote principal
package main

//importar função fmt
import "fmt"

//declaração da variável CONST do ponto de ebulição da água em Kelvin
const ebulicaoK = 373

//função principal
func main() {

	tempK := ebulicaoK // variável do valor da temperatura em Kelvin
	tempC := (tempK - 273) //Conversão de Kelvin para Celsius
	//mostrar o resultado
	fmt.Printf("A temperatura de ebulição da água em K = %v, temperatura de ebulição da água em °C = %v .", tempK, tempC)

	// A gente espera que apareça K = 373 e C = 100

}