package main

import "fmt"

func main() {

	// Dentro do Golang temos o tipo zero para cada tipo de dado
	// O tipo zero é o valor que uma variável recebe quando é declarada sem valor

	var texto string;
	var numeroInteiro int;
	var booleano bool;
	var numeroDecimal float64;
	var ponteiro *int;

	fmt.Printf("O valor zero de uma string é: %v\n", texto);
	fmt.Printf("O valor zero de um inteiro é: %v\n", numeroInteiro);
	fmt.Printf("O valor zero de um booleano é: %v\n", booleano);
	fmt.Printf("O valor zero de um número decimal é: %v\n", numeroDecimal);
	fmt.Printf("O valor zero de um ponteiro é: %v\n", ponteiro);
}