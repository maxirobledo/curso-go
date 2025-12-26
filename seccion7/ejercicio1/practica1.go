package main

import (
	"encoding/json"
	"fmt"

	"github.com/maxirobledo/curso-go/packages/fileops"
	"github.com/maxirobledo/curso-go/packages/products"
)

func main() {

	//1)
	hobbies := [3]string{"play football", "read books", "learn programming"}

	fmt.Println("Punto 1)")
	fmt.Println(hobbies)

	//2)
	fmt.Println("Punto 2)")
	fmt.Println("Primer hobbie: ", hobbies[0])
	newHobbies := hobbies[1:]
	fmt.Println("Nuevos hobbies: ", newHobbies)

	//3)
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println("Punto 3)")
	fmt.Println("Primera manera: ", slice1)
	fmt.Println("Segunda manera: ", slice2)

	//4)
	slice1[0] = hobbies[1]
	slice1[1] = hobbies[2]
	fmt.Println("Punto 4)")
	fmt.Println(slice1)

	//5)
	goals := []string{}
	goals = append(goals, "Learn the basics of Go", "Improve my skills")
	fmt.Println("Punto 5)")
	fmt.Println(goals)

	//6)
	goals[1] = "Getting started With Go Programming"
	goals = append(goals, "Start my new job with a foundaton in Go Programming")
	fmt.Println("6)")
	fmt.Println(goals)

	//Bonus
	prodcutsList := []*products.Product{}

	product1, err := products.New("Laptop", 1, 1500)
	if err != nil {
		fmt.Println("Error creating product:", err)
		return
	}

	product2, err := products.New("Iphone", 2, 1300)
	if err != nil {
		fmt.Println("Error creating product:", err)
		return
	}

	prodcutsList = append(prodcutsList, product1, product2)

	var streams string = ""
	// Mostrar los productos actuales
	for _, p := range prodcutsList {
		fmt.Println("#######################")
		fmt.Println(products.GetProduct(p))
		streams = streams + products.GetProduct(p)
	}

	// Convertir la lista de productos a JSON
	jsonData, err := json.MarshalIndent(prodcutsList, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Guardar datos en un file json
	fileops.WiteJsonToFile(jsonData, "products.json")
	println("Datos guardados en products.json")

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
