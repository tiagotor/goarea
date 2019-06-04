package goarea

import "math"

// Pi é uma proporção numérica definida pela relação entre
// o perímetro de uma circunferencia e seu dimâmetro.
const Pi = 3.141516

//Circ é responsável por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um quadrado
func Rect(base, altura float64) float64 {
	return base * altura
}

//não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
