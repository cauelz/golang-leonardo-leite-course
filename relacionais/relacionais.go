package main

import "fmt"


func main() {

	// comparação de igualdade
	//Exemplo: 1 == 1 => true
	fmt.Println("1 == 1: ", 1 == 1)
	// comparação de desigualdade
	//Exemplo: 1 != 1 => false
	fmt.Println("1 != 1: ", 1 != 1)
	// comparação de maior que
	//Exemplo: 1 > 1 => false
	fmt.Println("1 > 1: ", 1 > 1)
	// comparação de menor que
	//Exemplo: 1 < 1 => false
	fmt.Println("1 < 1: ", 1 < 1)
	// comparação de maior ou igual
	//Exemplo: 1 >= 1 => true
	fmt.Println("1 >= 1: ", 1 >= 1)
	// comparação de menor ou igual
	//Exemplo: 1 <= 1 => true
	fmt.Println("1 <= 1: ", 1 <= 1)

	// comparando datas

	// 01/01/2019
	data1 := "01/01/2019"
	// 01/01/2019
	data2 := "01/01/2019"

	// comparando se as datas são iguais
	fmt.Println("data1 == data2: ", data1 == data2)

	// Go permite comparar structs, mas apenas se todos os campos forem comparáveis
	type Pessoa struct {
		Nome string
		Idade int
	}

	pessoa1 := Pessoa{"João", 20}
	pessoa2 := Pessoa{"João", 20}


	// comparando se as pessoas são iguais
	fmt.Println("pessoa1 == pessoa2: ", pessoa1 == pessoa2)

	// comparando pessoas diferentes
	pessoa3 := Pessoa{"Maria", 20}
	fmt.Println("pessoa1 == pessoa3: ", pessoa1 == pessoa3)

	// se a struct contiver campos não comparáveis, o código não compilará
	type Pedido struct {
		Numero int
		Cliente Pessoa
		Items []string
	}

	// pedido1 := Pedido{1, pessoa1, []string{"item1", "item2"}}
	// pedido2 := Pedido{1, pessoa1, []string{"item1", "item2"}}

	// comparando se os pedidos são iguais
	// fmt.Println("pedido1 == pedido2: ", pedido1 == pedido2)
	// Error: invalid operation: pedido1 == pedido2 (struct containing Pessoa cannot be compared)

}