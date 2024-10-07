package exercice

import (
	"fmt"
	"math"
)

func formuleDAntoine(A float64, B float64, C float64, t float64) float64 {
	return math.Pow(10, A-(B/(t+C)))
}

func calculYRaoulDalton(tVap float64, x float64, ptot float64) float64 {
	return (tVap * x) / ptot
}

func n3Calcul(y3 float64, n1 float64, x4 float64) float64 {
	return (n1 - (x4 * n1)) / (y3 - x4)
}

func Exercice2() {

	var x4hexa float64 = 0.817
	var pTotAtm float64 = 0.854
	// var n1 float64 = 3500

	fmt.Println("Hello from exercice2")
	p0pent := formuleDAntoine(6.84471, 1060.793, 231.541, 55)
	p0hexa := formuleDAntoine(6.88551, 1175.817, 224.867, 55)
	y3nhexane := calculYRaoulDalton(p0hexa, x4hexa, pTotAtm*760.0)
	y3nPentane := calculYRaoulDalton(p0pent, (1 - x4hexa), pTotAtm*760.0)

	fmt.Printf("Tention vapeur hexane: %f \n", p0hexa)
	fmt.Printf("Tention vapeur pentane: %f \n", p0pent)
	fmt.Printf("y3nhexane : %f \n", y3nhexane)
	fmt.Printf("y3nPentane : %f \n", y3nPentane)
	// fmt.Printf("n3 : %f \n", n3)
}
