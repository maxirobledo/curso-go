// Calculadora de Beneficios en Go
// El usuario ingresa los ingresos, los gastos y la tasa de impuestos
// Calcular Beneficios antes de impuesto (EBT), beneficios después de impuesto (EAT) y relacion entre EBT y EAT.

package main

import "fmt"

func main() {
	var ingresos, gastos, tasaImpuestos float64

	fmt.Print("Ingrese sus ingresos: ")
	fmt.Scan(&ingresos)
	fmt.Print("Ingrese sus gastos: ")
	fmt.Scan(&gastos)
	print("Ingrese la tasa de impuestos: ")
	fmt.Scan(&tasaImpuestos)

	ebt := ingresos - gastos
	eat := ebt * (1 - tasaImpuestos/100)
	ratio := ebt / eat
	fmt.Printf("Ingresos: %.2f\nGastos: %.2f\nEBT: %.2f\nEAT: %.2f\nRelacion entre EBT y EAT: %.2f\n", ingresos, gastos, ebt, eat, ratio)
}
