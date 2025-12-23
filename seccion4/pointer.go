package main

import "fmt"

func getAdultAge(p *int) {
	*p = *p - 18
}

func main() {
	age := 37

	var agePtr *int = &age

	fmt.Printf("Age pointer is: %d\n", *agePtr)

	getAdultAge(agePtr)
	fmt.Printf("Age is: %d\n", age)
}
