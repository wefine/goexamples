package main

import "fmt"

type Team struct {
	Name, Description string
	Users             []User
}

func (t Team) GetTeamDetails() {
	fmt.Printf("Team: %s  - %s\n", t.Name, t.Description)
	fmt.Println("Details   of the team members:")
	for _, v := range t.Users {
		v.PrintName()
		v.PrintDetails()
	}
}
