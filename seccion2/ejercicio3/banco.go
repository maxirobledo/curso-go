package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const file = "saldos.txt"

func leerArchivo() (float64, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return 0, errors.New("No se encontro el archivo de saldo.")
	}
	saldo, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("No es posible recuperar el saldo.")
	}
	return saldo, nil
}

func escribirArchivo(saldo float64) {
	saldoTexto := fmt.Sprint(saldo)
	err := os.WriteFile(file, []byte(saldoTexto), 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
		return
	}
	fmt.Println("Saldo guardado en saldos.txt")
}

func main() {

	var saldo, err = leerArchivo()

	if err != nil {
		fmt.Println("ERROR - ", err)
		panic("La aplicacion se cerrara.")
	}
	var opcion int

	fmt.Println("Bienvenido al Banco Go")

	for {

		fmt.Println("===================================")
		fmt.Println("Seleccione una opcion:")
		fmt.Println("1. Consultar saldo")
		fmt.Println("2. Depositar dinero")
		fmt.Println("3. Retirar dinero")
		fmt.Println("4. Salir")
		fmt.Print("Ingrese su opcion: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			fmt.Printf("Su saldo actual es: %.2f\n", saldo)

		case 2:
			var deposito float64
			fmt.Print("Ingrese el monto a depositar: ")
			fmt.Scan(&deposito)
			if deposito <= 0 {
				fmt.Println("El monto a depositar debe ser mayor que cero.")
				continue
			}
			saldo += deposito
			fmt.Printf("Deposito exitoso!!!\nSu nuevo saldo es: %.2f\n", saldo)
			escribirArchivo(saldo)
		case 3:
			var retiro float64
			fmt.Print("Ingrese el monto a retirar: ")
			fmt.Scan(&retiro)
			if retiro <= 0 {
				fmt.Println("El monto a retirar debe ser mayor que cero.")
				continue
			}
			if retiro > saldo {
				fmt.Println("Fondos insuficientes para realizar el retiro.")
				continue
			}
			saldo -= retiro
			fmt.Printf("Retiro exitoso!!!.\nSu nuevo saldo es: %.2f\n", saldo)
			escribirArchivo(saldo)
		default:
			fmt.Printf("Gracias por usar el Banco Go. ¡Hasta luego!\n")
			return
		}
	}

}
