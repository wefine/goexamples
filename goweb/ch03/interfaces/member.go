package main

import "fmt"

type Member struct {
	Person //type embedding for composition
	Skills []string
}

//overrides PrintDetails
func (m Member) PrintDetails() {
	//Call person PrintDetails
	m.Person.PrintDetails()
	fmt.Println("Skills:")
	for _, v := range m.Skills {
		fmt.Println(v)
	}
}
