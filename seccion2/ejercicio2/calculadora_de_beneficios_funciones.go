// Calculadora de Beneficios en Go
// El usuario ingresa los ingresos, los gastos y la tasa de impuestos
// Calcular Beneficios antes de impuesto (EBT), beneficios después de impuesto (EAT) y relacion entre EBT y EAT.

package main

import "fmt"

func main() {

	ingresos := pedirDatos("Ingrese los ingresos: ")
	gastos := pedirDatos("Ingrese los gastos: ")
	tasaImpuestos := pedirDatos("Ingrese la tasa de impuestos (%): ")

	ebt, eat, ratio := calcularBeneficios(ingresos, gastos, tasaImpuestos)

	fmt.Printf("Ingresos: %.2f\nGastos: %.2f\nEBT: %.2f\nEAT: %.2f\nRelacion entre EBT y EAT: %.2f\n", ingresos, gastos, ebt, eat, ratio)
}

func pedirDatos(mensaje string) float64 {
	var data float64
	fmt.Print(mensaje)
	fmt.Scan(&data)
	return data
}

func calcularBeneficios(ingresos, gastos, tasaImpuestos float64) (float64, float64, float64) {
	ebt := ingresos - gastos
	eat := ebt * (1 - tasaImpuestos/100)
	ratio := ebt / eat
	return ebt, eat, ratio
}
