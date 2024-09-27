package exercice

import (
	"fmt"
	"math"
)

var ratioN2N3 float64 = 1.5
var n3 float64 = 120.0
var temps float64 = 40.0
var y3b float64 = 0.94
var y3p float64 = 1 - y3b
var (
	n2   float64
	p0B  float64
	p0P  float64
	Ptot float64
)

// func calculPoAvecLoiAntoine() float64 {
//
// }

func formuleDAntoine(A float64, B float64, C float64, t float64) float64 {
	return math.Pow(10, A-B/(t+C))
}

func calculN2() float64 {
	n2 = n3 * ratioN2N3
	return n2
}

func calculPtot() float64 {
	Ptot = p0B*y3b + p0P*y3p
	return Ptot
}

func Exercice1() {
	fmt.Println("Exercice1")
	fmt.Printf("Valeur n2: %f \n", calculN2())

	p0B = formuleDAntoine(6.82272, 943.453, 239.888, temps)
	p0P = formuleDAntoine(6.84471, 1060.793, 231.541, temps)
	fmt.Printf("Valeur PoB: %f \n", p0B)
	fmt.Printf("Valeur PoP: %f \n", p0P)
	fmt.Printf("Valeur de Ptot: %f \n", calculPtot())

}
