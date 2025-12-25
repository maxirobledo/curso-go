// app que le pide datos de entrada al usuario, y crea un archivo json con esos datos usando structs y packages

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/maxirobledo/curso-go/packages/fileops"
	"github.com/maxirobledo/curso-go/packages/note"
)

func main() {

	titulo := getData("Ingrese el título de la informacion: ")

	contenido := getData("Ingrese el contenido de la informacion: ")

	var nota1 *note.Note
	nota1, err := note.New(titulo, contenido)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Título: %s\nContenido: %s\nCreado el: %s\n", nota1.Title, nota1.Content, nota1.CreatedAt)

	err = guardarJSONEnArchivo(nota1)
	if err != nil {
		fmt.Println("Error al guardar la nota en archivo:", err)
		return
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

func guardarJSONEnArchivo(nota1 *note.Note) error {
	fileName := nota1.Title
	fileName = strings.ReplaceAll(fileName, " ", "_") + ".json"
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(nota1)

	if err != nil {
		return err
	}

	fileops.WiteJsonToFile(json, fileName)
	return nil
}
