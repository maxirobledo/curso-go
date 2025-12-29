package main

import "fmt"

func main() {
	fact := factorial(6)
	fmt.Println(fact)
}

func factorial(numero int) int {
	if numero == 0 {
		return 1
	}

	return numero * factorial(numero-1)
	// resultado := 1

	// for i := 1; i <= numero; i++ {
	// 	resultado *= i
	// }

	// return resultado
}
