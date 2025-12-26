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
	"github.com/maxirobledo/curso-go/packages/todo"
)

type saver interface {
	Save() error
}

func main() {

	// titulo := getData("Ingrese el título de la informacion: ")
	// contenido := getData("Ingrese el contenido de la informacion: ")
	text := getData("Ingrese el texto de la tarea pendiente: ")

	task, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	task.Display()

	saveData(task)

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Datos guardados correctamente")
	return nil
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
