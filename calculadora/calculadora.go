package calculadora

import "fmt"

type Calculadora interface {
	Soma(int, int)
	Subtracao(int, int)
	Divisao(int, int)
	GetNumero() int
	SetNumero(int)
}

func NewCalculadora() Calculadora {
	return calculadora{}
}

type calculadora struct {
	NumeroResultado int
}

func (c calculadora) Soma(x, y int) {
	c.SetNumero(x + y)
}

func (c calculadora) Subtracao(x, y int) {
	c.SetNumero(x - y)
}

func (c calculadora) Divisao(x, y int) {
	c.SetNumero(x / y)
}

func (c calculadora) GetNumero() int {
	fmt.Print(c.NumeroResultado)
	return (c.NumeroResultado)
}

func (c calculadora) SetNumero(numero int) {
	c.NumeroResultado = numero

}
