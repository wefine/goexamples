package main

import "fmt"

// Declare constant
const Title = "Person Details"

// Declare package variable
var Country = "USA"

func main() {
	// Print constant variable
	fmt.Println(Title)

	fname, lname := "Shiju", "Varghese"
	age := 35

	// Print local variables
	fmt.Println("First Name:", fname)
	fmt.Println("Last Name:", lname)
	fmt.Println("Age:", age)

	// Print package variable
	fmt.Println("Country:", Country)
}
