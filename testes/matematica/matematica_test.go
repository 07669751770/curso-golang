package matematica

import (
	"fmt"
	"testing"
)

func TestMedia(t *testing.T) {
	//t.Parallel()

	media := Media(7.7, 8.1, 5.9, 10.0)
	esperado := 0.0
	esperado += 7.7
	esperado += 8.1
	esperado += 5.9
	esperado += 10.0
	esperado /= float64(4)

	fmt.Println("Média:", media)
	fmt.Println("Esperado:", esperado)

	if media != esperado {
		t.Errorf("Valor esperado %f, mas o resultado encontrado foi %f", esperado, media)
	}
}
