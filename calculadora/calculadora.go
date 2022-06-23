package calculadora

type Calculadora interface {
	Soma(int, int) int
	Subtracao(int, int) int
	Divisao(int, int) int
}

func NewCalculadora() Calculadora {
	return calculadora{}
}

type calculadora struct {
}

func (c calculadora) Soma(x, y int) int {
	return (x + y)
}

func (c calculadora) Subtracao(x, y int) int {
	return (x - y)
}

func (c calculadora) Divisao(x, y int) int {
	return (x / y)
}
