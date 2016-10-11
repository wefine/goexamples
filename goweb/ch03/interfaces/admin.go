package main

import "fmt"

type Admin struct {
	Person //type embedding for composition
	Roles []string
}

//overrides PrintDetails
func (a Admin) PrintDetails() {
	//Call person PrintDetails
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}
