// Calculadora de Beneficios en Go
// El usuario ingresa los ingresos, los gastos y la tasa de impuestos
// Calcular Beneficios antes de impuesto (EBT), beneficios después de impuesto (EAT) y relacion entre EBT y EAT.
// 1. Validar entrada del usuario
//	 a. Ingresos y gastos deben ser números positivos y mayores a cero
//	 b. Mostrar mensaje de error si la entrada no es válida y exit de la aplicación
// 2. Almacenar los datos de resultados en un archivo

package main

import (
	"errors"
	"fmt"
	"os"
)

func pedirDatos(mensaje string) (float64, error) {
	var data float64
	fmt.Print(mensaje)
	fmt.Scan(&data)
	if data <= 0 {
		return 0, errors.New("ERROR - El valor debe ser mayor que cero.")
	}
	return data, nil
}

func calcularBeneficios(ingresos, gastos, tasaImpuestos float64) (float64, float64, float64) {
	ebt := ingresos - gastos
	eat := ebt * (1 - tasaImpuestos/100)
	ratio := ebt / eat
	return ebt, eat, ratio
}

func escribirArchivo(resultado string) error {
	err := os.WriteFile("beneficios.txt", []byte(resultado), 0644)
	if err != nil {
		return errors.New("ERROR - No es posible escribir el archivo.")
	}
	fmt.Println("Se guardaron los resultados en el archivo beneficios.txt")
	return nil
}

func main() {

	ingresos, err1 := pedirDatos("Ingrese los ingresos: ")
	gastos, err2 := pedirDatos("Ingrese los gastos: ")
	tasaImpuestos, err3 := pedirDatos("Ingrese la tasa de impuestos (%): ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("ERROR - Entrada no valida. Asegurese de ingresar numeros positivos mayores a cero.")
		panic("La aplicacion se cerrara.")
	}

	ebt, eat, ratio := calcularBeneficios(ingresos, gastos, tasaImpuestos)
	resultado := fmt.Sprintf("Ingresos: %.2f\nGastos: %.2f\nEBT: %.2f\nEAT: %.2f\nEBT/EAT: %.2f\n", ingresos, gastos, ebt, eat, ratio)
	err := escribirArchivo(resultado)
	if err != nil {
		fmt.Println("ERROR al escribir el archivo - ", err)
		panic("La aplicacion se cerrara.")
	}

	fmt.Printf("Ingresos: %.2f\nGastos: %.2f\nEBT: %.2f\nEAT: %.2f\nRelacion entre EBT y EAT: %.2f\n", ingresos, gastos, ebt, eat, ratio)
}
