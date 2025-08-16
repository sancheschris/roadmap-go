package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

// this is not exported - you can only use it inside this package (use capital letter to export)
func dividir(a, b float64) float64 {
	if (b > 0) {
		return a / b
	}
	return 0.0
}