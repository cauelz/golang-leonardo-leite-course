package main

import (
	"fmt"
	"math"
)


func main() {
	
	a := 10
	b := 5

	fmt.Println("------------ Operações aritméticas ------------")

	// soma
	fmt.Println("Adição: ", a + b)

	// subtração
	fmt.Println("Subtração: ", a - b)

	// divisão
	fmt.Println("Divisão: ", a / b)

	// multiplicação
	fmt.Println("Multiplicação: ", a * b)

	// resto da divisão ou módulo
	fmt.Println("Resto da divisão/modulo: ", a % b)

	// incremento
	a++
	fmt.Println("Incremento: ", a)

	// decremento
	b--
	fmt.Println("Decremento: ", b)

	// operações bitwise
	fmt.Println("------------ Operações bitwise ------------")
	// O que são operações bitwise?
	// São operações que são feitas bit a bit, ou seja, cada bit é comparado com o outro bit da mesma posição
	// Exemplo:
	// 10 em binário é 1010
	// 5 em binário é 0101

	//Operação AND: Compara bit a bit e retorna 1 se ambos os bits forem 1
	fmt.Println("10 & 5: ", 10 & 5) // 1010 & 0101 = 0000

	//Operação OR: Compara bit a bit e retorna 1 se pelo menos um dos bits for 1
	fmt.Println("10 | 5: ", 10 | 5) // 1010 | 0101 = 1111

	//Operação XOR: Compara bit a bit e retorna 1 se os bits forem diferentes
	fmt.Println("10 ^ 5: ", 10 ^ 5) // 1010 ^ 0101 = 1111

	//Operação NOT: Inverte os bits
	fmt.Println("^10: ", ^10) // ^1010 = 0101

	//Operação deslocamento para a esquerda: Desloca os bits para a esquerda
	fmt.Println("10 << 1: ", 10 << 1) // 1010 << 1 = 10100

	//Operação deslocamento para a direita: Desloca os bits para a direita
	fmt.Println("10 >> 1: ", 10 >> 1) // 1010 >> 1 = 0101

	fmt.Println("------------ Utilizando o pacote Math ------------")
	
	// Para utilizar o pacote math, é necessário importá-lo
	// import "math"
	// O pacote math possui funções matemáticas prontas para serem utilizadas

	//Constantes
	fmt.Println("Pi: ", math.Pi)
	fmt.Println("Euler: ", math.E)

	//Funções matemáticas
	fmt.Println("Raiz quadrada de 16: ", math.Sqrt(16))
	fmt.Println("Potência de 2 elevado a 3: ", math.Pow(2, 3))
	fmt.Println("Arredonda para cima: ", math.Ceil(3.14))
	fmt.Println("Arredonda para baixo: ", math.Floor(3.14))
	fmt.Println("Arredonda para o inteiro mais próximo: ", math.Round(3.14))

}