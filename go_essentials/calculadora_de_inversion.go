package main

import (
	"fmt"
	"math"
)

func main() {
	const tasaInflacion, tasaRendimiento = 2.5, 5.5
	var montoInicial, aniosInversion float64

	fmt.Print("Ingrese el monto que desea invertir: ")
	fmt.Scan(&montoInicial)
	fmt.Print("Ingrese la cantidad de años que desea invertir: ")
	fmt.Scan(&aniosInversion)

	montoFinal := montoInicial * math.Pow((1+tasaRendimiento/100), aniosInversion)

	montoFinalAjustado := montoFinal / math.Pow((1+tasaInflacion/100), aniosInversion)

	fmt.Printf("Inversión inicial de %.2f\nAños de inversion %.0f\nTasa de inflacion %.2f\nTasa de rendimiento %.2f\nMonto Final Ajustado %.2f", montoInicial, aniosInversion, tasaInflacion, tasaRendimiento, montoFinalAjustado)
}
