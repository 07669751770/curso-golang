package main

import "fmt"

func main() {
	var notas [3]float64
	fmt.Println(notas) // [0 0 0]

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas) // [7.8 4.3 9.1]

	total := 0.0

	for nota := range notas {
		total += notas[nota]
	}
	fmt.Println(total) // 21.2

	media := total / float64(len(notas))
	fmt.Println(media) // 7.066666666666666
}
