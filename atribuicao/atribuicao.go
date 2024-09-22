package main

import "fmt"


func main() {

	// byte é um alias para uint8
	var a byte = 10
	fmt.Println("byte a: ", a)

	// rune é um alias para int32. Atribuimos um valor unicode e retorna o valor decimal
	var b rune = 'a'
	fmt.Println("rune b: ", b)

	// inferencia de tipo
	c := 10
	
	c += 10 // c = c + 10
	fmt.Println("c += 10: ", c)

	c -= 10 // c = c - 10
	fmt.Println("c -= 10: ", c)

	c *= 10 // c = c * 10
	fmt.Println("c *= 10: ", c)

	c /= 10 // c = c / 10
	fmt.Println("c /= 10: ", c)

	c %= 10 // c = c % 10
	fmt.Println("c %= 10: ", c)

	// inferencia de multiplas variaveis
	d, e := 10, 5
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)

}