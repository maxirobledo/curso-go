package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("ERROR - Los campos título y contenido son requeridos")
	}

	return Todo{
		Text: content,
	}, nil

}

func (todo Todo) Display() {
	println("Todo:", todo.Text)
}

func (todo Todo) Save() error {
	// Lógica para guardar la tarea pendiente
	filename := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}
