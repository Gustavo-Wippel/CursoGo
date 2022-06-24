package calculadora

import (
	"testing"
)

var (
	calc calculadora = calculadora{}
)

func TestSome(t *testing.T) {
	resultado := //calc.Soma(3, 2)
		5
	if resultado != 4 {
		t.Errorf("O resultado esperado foi diferente de 5: resultado")
	}
}
