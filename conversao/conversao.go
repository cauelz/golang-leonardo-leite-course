package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Neste primeiro exemplo, vemos que para converter um texto que representa um numero decimal, em decimal
	// precisamos utilizar o pacote strconv e acessar a funcão ParseFloat(s string, bit int)
	// e depois, para converter o decimal em inteiro, usamos a declaração direta int(decimal)
	var textoNumero = "4.28"

	decimal, _ := strconv.ParseFloat(textoNumero, 64)
	inteiro:= int(decimal)
	fmt.Println("Convertendo para decimal: ", decimal)
	fmt.Println("Convertendo para inteiro: ", inteiro)

	// Podemos converter um inteiro para texto
	numero := 5
	convertido := strconv.Itoa(numero)
	fmt.Println("Inteiro para texto: " + convertido)

	// Se quisermos converter qualquer numero para texto, temos que tomar cuidado
	// a declaracao direta string(numero) vai converter para unicode.
	fmt.Println("Convertendo numero para texo (retorna unicode): " + string(55))

	// inteiro para texto
	fmt.Println("Inteiro para texto: " + strconv.Itoa(inteiro))

	// decimal para texto
	fmt.Println("Decimal para texto: " + strconv.FormatFloat(decimal, 'f', -1, 64))



}