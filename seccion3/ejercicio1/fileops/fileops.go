package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("No se encontro el archivo.")
	}
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("No es posible recuperar el valor.")
	}
	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
		return
	}
	fmt.Println("Valor escrito correctamente en el archivo.")
}
