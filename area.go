package area

import "math"

//PI public
const PI = 3.1415

func Circ(raio float64) float64 {
	return PI * math.Pow(raio, 2)
}

func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visível
func _TriangulaEq(base, altura float64) float64 {
	return base * altura / 2
}
