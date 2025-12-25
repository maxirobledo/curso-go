package main

import (
	"fmt"

	"github.com/maxirobledo/curso-go/packages/fileops"
)

const file = "saldos.txt"

func main() {

	var saldo, err = fileops.GetFloatFromFile(file)

	if err != nil {
		fmt.Println("ERROR - ", err)
		panic("La aplicacion se cerrara.")
	}
	var opcion int

	fmt.Println("Bienvenido al Banco Go")

	for {

		mostrarMenu()
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
			fileops.WriteFloatToFile(saldo, file)
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
			fileops.WriteFloatToFile(saldo, file)
		default:
			fmt.Printf("Gracias por usar el Banco Go. ¡Hasta luego!\n")
			return
		}
	}

}
