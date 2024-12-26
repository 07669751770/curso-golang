package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	//t.Parallel()

	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Teste", "e", 1},
		{"Teste", "T", 0},
		{"Teste", "z", -1},
		{"", "", 0},
		{"Opa", "pa", 1},
	}

	for _, teste := range testes {
		t.Logf("Massa de teste: %v", teste)
		retorno := strings.Index(teste.texto, teste.parte)
		if retorno != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, retorno)
		}
	}

}
