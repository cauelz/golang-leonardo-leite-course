package main

import (
	"fmt"
	"math"
)


func main() {
	// Variaveis em Golang
	// var nomeDaVariavel tipoDaVariavel = valorDaVariavel
	var PI float64 = 3.14159265359
	// ou
	// var nomeDaVariavel = valorDaVariavel. O compilador infere o tipo da variavel
	var raio = 3.2

	// ou
	// Forma reduzida de declarar variaveis
	// nomeDaVariavel := valorDaVariavel
	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	// Podemos declarar mais uma variavel na forma reduzida com diferentes tipos
	// var x, y, z, w := 1, 2, true, "opa"

	// Constantes
	// Não precisa separar com virgula
	const (
		a = 1
		b = 2
		c = 3
	)

	fmt.Println(a, b, c)
}