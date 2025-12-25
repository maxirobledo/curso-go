package user

import (
	"errors"
	"fmt"
)

type User struct {
	nombre   string
	apellido string
	edad     string
}

func New(nombre string, apellido string, edad string) (*User, error) {
	if nombre == "" || apellido == "" || edad == "" {
		return nil, errors.New("ERROR - todos los campos son requeridos")
	}

	return &User{
		nombre:   nombre,
		apellido: apellido,
		edad:     edad,
	}, nil
}

func (u User) OutputUserData() {
	fmt.Printf("Nombre: %s\n", u.nombre)
	fmt.Printf("Apellido: %s\n", u.apellido)
	fmt.Printf("Edad: %s\n", u.edad)
}

func (u *User) ClearUserData() {
	u.nombre = ""
	u.apellido = ""
	u.edad = ""
}

func (u *User) GetUserData() string {
	return fmt.Sprintf("Nombre: %s\nApellido: %s\nEdad: %s\n---------------------\n", u.nombre, u.apellido, u.edad)
}
