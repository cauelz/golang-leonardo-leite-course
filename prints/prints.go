package main

import (
	"fmt"
	"strconv"
)


func main() {

	// Print não pula linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// Println pula linha
	fmt.Println(" Nova")
	fmt.Println("linha.")
	fmt.Print("Nova")

	// Brincando com strings e tipos de Prints
	x := 3.141516

	// Sprintf retorna uma string formatada.
	// %.2f é um formatador de float com duas casas decimais.
	xStr := fmt.Sprintf("%.2f", x)
	// Println imprime a string.
	fmt.Println("O valor de x é", xStr)

	// FormatFloat retorna uma string formatada.
	// 'f' é o tipo de formatação.
	// 2 é o número de casas decimais.
	// 64 é o bitSize.
	xStr = strconv.FormatFloat(x, 'f', 2, 64)
	fmt.Println("O valor de x é", xStr)

	//Printf imprime a string formatada.
	// %.2f é um formatador de float com duas casas decimais.
	fmt.Printf("O valor de x é %.2f\n", x)

	// Printf com varios tipos de formatação.
	a := 1
	b := 1.999999
	c := false
	d := "opa"
	// %d é um formatador de int.
	// %f é um formatador de float.
	// %1.f é um formatador de float com uma casa decimal.
	// %t é um formatador de bool.
	// %s é um formatador de string.
	// %v é um formatador de qualquer tipo.
	fmt.Printf("%d, %f, %1.f, %t, %s\n", a, b, b, c, d)
	fmt.Printf("%v, %v, %.1v, %v, %v\n", a, b, b, c, d)
}