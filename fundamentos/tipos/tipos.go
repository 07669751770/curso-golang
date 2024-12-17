package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// número iteiros
	var a int = 1
	fmt.Println(a)
	fmt.Println("O literal inteiro é", reflect.TypeOf(a))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal ... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // Represemta um mapeamento para tabela UNICODE (int32)
	fmt.Println("O tipo de i2 é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal é", reflect.TypeOf(49.99)) // float64

	// Boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string

	s1 := "Olá meu nome é José"
	fmt.Println(s1 + "!")

	fmt.Println("O tamanho da string é", len(s1))

	// string com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	José`

	fmt.Println("O tamanho da string é", len(s2))

	// char???
	caractere := 'a'
	fmt.Println("O tipo de caractere é", reflect.TypeOf(caractere))
	fmt.Println(caractere)

}
