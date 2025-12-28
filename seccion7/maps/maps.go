package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Mapa que asocia nombres de países con sus capitales
	capitales := map[string]string{
		"España":   "Madrid",
		"Francia":  "París",
		"Alemania": "Berlín",
		"Italia":   "Roma",
		"Portugal": "Lisboa",
	}

	// Imprimir las capitales de los países
	for pais, capital := range capitales {
		println(capital + " es la capital de " + pais)
	}

	fmt.Println("------------------------")
	pais := getData("Ingrese el pais del cual quiere consultar capital: ")
	capital, existe := capitales[pais]
	if existe {
		fmt.Printf("La capital de %s es %s\n", pais, capital)
	} else {
		fmt.Printf("No se encontró la capital para el país: %s\n", pais)
	}

	fmt.Println("------------------------")
	// Agregar un nuevo país y su capital
	nuevoPais := getData("Ingrese el nuevo país que desea agregar: ")
	nuevaCapital := getData("Ingrese la capital de " + nuevoPais + ": ")
	capitales[nuevoPais] = nuevaCapital

	fmt.Println("------------------------")
	fmt.Println("Capitales actualizadas:")
	for pais, capital := range capitales {
		println(capital + " es la capital de " + pais)
	}

	fmt.Println("------------------------")
	// Eliminar un país del mapa
	paisAEliminar := getData("Ingrese el país que desea eliminar: ")
	_, existe = capitales[paisAEliminar]
	if existe {
		delete(capitales, paisAEliminar)
		fmt.Printf("Se eliminó %s del mapa.\n", paisAEliminar)
	} else {
		fmt.Printf("No se encontró %s en el mapa.\n", paisAEliminar)
	}

}

func getData(prompt string) string {
	print(prompt)
	reader := bufio.NewReader(os.Stdin)
	texto, error := reader.ReadString('\n')
	if error != nil {
		fmt.Println("Error leyendo la entrada:", error)
		return ""
	}
	// Eliminar el salto de línea
	texto = strings.TrimSuffix(texto, "\n")
	texto = strings.TrimSuffix(texto, "\r")
	return texto
}
