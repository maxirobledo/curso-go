package main

import (
	"fmt"
)

func main() {

	rankingCursos := make(map[string]float64, 3)

	rankingCursos["Go"] = 4.9
	rankingCursos["Python"] = 4.8
	rankingCursos["JavaScript"] = 4.7
	rankingCursos["Java"] = 4.5

	for curso, ranking := range rankingCursos {
		println("El curso de " + curso + " tiene un ranking de " + fmt.Sprintf("%.1f", ranking))
	}
}
