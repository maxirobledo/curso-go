package main

import (
	"fmt"

	"github.com/maxirobledo/curso-go/packages/fileops"
	"github.com/maxirobledo/curso-go/packages/user"
)

func getData(texto string) string {
	fmt.Print(texto)
	var valor string
	fmt.Scanln(&valor)
	return valor
}

func main() {

	nombreUsuario := getData("Ingrese su nombre: ")
	apellidoUsuario := getData("Ingrese su apellido: ")
	edadUsuario := getData("Ingrese su edad: ")

	var u1 *user.User

	u1, err := user.New(nombreUsuario, apellidoUsuario, edadUsuario)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileops.WriteStructToFile(u1.GetUserData(), "user.txt")
	u1.OutputUserData()
}
