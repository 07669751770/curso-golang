package main

import "fmt"

func main() {

	nota := 4

	switch nota {
	case 10:
		fmt.Println("Aprovado com louvor!")
		break
	case 9:
		fmt.Println("Aprovado com mérito!")
		break
	case 8, 7:
		fmt.Println("Aprovado!")
		break
	case 6, 5:
		fmt.Println("Recuperação!")
		break
	case 4, 3:
		fmt.Println("Recuperação com dificuldades!")
		fallthrough // executa o próximo case
	case 2, 1, 0:
		fmt.Println("Reprovado!")
		break
	default:
		fmt.Println("Nota inválida!")
		break
	}

}
