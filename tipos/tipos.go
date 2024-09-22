package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// numeros inteiros
	fmt.Println("1 + 1 =", reflect.TypeOf(1 + 1))
	fmt.Println("255", reflect.TypeOf(255))

	// numeros inteiros em bytes SEM sinal (uint8, uint16, uint32, uint64)
	var a uint8 = 255
	fmt.Println("a uint8 é", reflect.TypeOf(a))

	// tipo byte é um alias para uint8
	b := byte(255)
	fmt.Println("b byte é", reflect.TypeOf(b))

	// numeros inteiro em bytes COM sinal (int8, int16, int32, int64)
	c := int(-255)
	fmt.Println("c int é", reflect.TypeOf(c))
	d := math.MaxInt64
	fmt.Println("d int é", reflect.TypeOf(d))
	fmt.Println("O valor máximo de d é", d)

	// numero que representa um caractere unicode (rune)
	var e rune = 'a'
	fmt.Println("e rune é", reflect.TypeOf(e))
	fmt.Println("O valor de e é", e)

	// numero que representa um char (int32), mas ele não é um tipo existente.
	// veja que tivemos que usar aspas simples.
	var f1 = 'a'
	fmt.Println("f1 int32 é", reflect.TypeOf(f1))
	
	// numeros reais (float32, float64)
	f := float32(49.99)
	fmt.Println("f float32 é", reflect.TypeOf(f))
	g := float64(49.99)
	fmt.Println("g float64 é", reflect.TypeOf(g))
	f = 49.99
	fmt.Println("g float64 é", reflect.TypeOf(g))

	// boolean
	h := true
	fmt.Println("h bool é", reflect.TypeOf(h))
	// negação
	fmt.Println("O valor de h é", !h)

	// string
	i := "Olá, meu nome é Caue"
	fmt.Println("i string é", reflect.TypeOf(i))
	// tamanho da string
	fmt.Println("O tamanho de i é", len(i))
	// string com multiplas linhas
	i = `Olá
	meu
	nome
	é
	Caue`
	fmt.Println("i string é", reflect.TypeOf(i))
	fmt.Println("O tamanho de i é", len(i))

}