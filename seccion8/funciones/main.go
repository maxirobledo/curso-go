package main

// Ejemplo de uso de funciones como parámetros
// La funcion generica transforNumbers recibe un slice de enteros y una función que transforma cada entero. Devuelve un nuevo slice con los enteros transformados.
// Se definen dos funciones de transformación: double y triple, que duplican y triplican un entero respectivamente.

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubled := transforNumbers(&numbers, double)
	tripled := transforNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transforNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, n := range *numbers {
		dNumbers = append(dNumbers, transform(n))
	}
	return dNumbers
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}
