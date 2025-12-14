// Calculadora de Inversión en Go
// Este programa calcula el monto final de una inversión después de un cierto número de años, ajustado por la inflación.
// El usuario ingresa el monto inicial, los años de inversión.
// Calcular Beneficios antes de impuesto (EBT), beneficios después de impuesto (EAT) y relacion entre EBT y EAT.

package main

import (
	"fmt"
	"math"
)

func main() {
	const tasaInflacion = 2.5
	var montoInicial, aniosInversion, tasaRendimiento float64

	fmt.Print("Ingrese el monto que desea invertir: ")
	fmt.Scan(&montoInicial)
	fmt.Print("Ingrese la cantidad de años que desea invertir: ")
	fmt.Scan(&aniosInversion)
	fmt.Print("Ingrese la tasa de rendimiento deseada: ")
	fmt.Scan(&tasaRendimiento)

	montoFinal := montoInicial * math.Pow((1+tasaRendimiento/100), aniosInversion)

	montoFinalAjustado := montoFinal / math.Pow((1+tasaInflacion/100), aniosInversion)

	ratio := montoFinal / montoFinalAjustado

	fmt.Printf("Inversión inicial de %.2f\nAños de inversion %.0f\nTasa de inflacion %.2f\nTasa de rendimiento %.2f\nBeneficios antes de impuesto %.2f\nBeneficios con impuestos %.2f\nRelacion entre beneficios %.2f", montoInicial, aniosInversion, tasaInflacion, tasaRendimiento, montoFinal, montoFinalAjustado, ratio)
}
