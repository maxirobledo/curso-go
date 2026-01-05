package fileops

import (
	"bufio"
	"encoding/json"
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

// Function to append a struct's data to a file
func WriteStructToFile(data string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
		return
	}

	fmt.Println("Estructura agregada correctamente al archivo.")
}

func WiteJsonToFile(jsonData []byte, fileName string) {
	os.WriteFile(fileName, jsonData, 0644)
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("No se pudo abrir el archivo")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error al leer del archivo")
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("No se pudo crear el archivo")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("No se pudo transformar el archivo data a JSON")
	}

	file.Close()
	return nil
}
