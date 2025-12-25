package note

import (
	"errors"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {

	if title == "" || content == "" {
		return nil, errors.New("ERROR - Los campos título y contenido son requeridos")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}
