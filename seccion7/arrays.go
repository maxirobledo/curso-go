package main

import "fmt"

func main() {
	prices := [4]float64{10.5, 20.0, 15.75, 30.0}

	newPrrices := prices[1:3]

	fmt.Println("Prices:", prices)

	fmt.Println("Valor de la posicion numero 3:", prices[2])

	fmt.Println("New Prices:", newPrrices)
}
