package main

import (
	"fmt"

	"../strutils"
)

func main() {
	str1, str2 := "Golang", "gopher"
	// Convert to upper case
	fmt.Println("To Upper Case:", strutils.ToUpperCase(str1))

	// Convert to lower case
	fmt.Println("To Lower Case:", strutils.ToUpperCase(str1))

	// Convert first letter to upper case
	fmt.Println("To First Upper:", strutils.ToFirstUpper(str2))
}
