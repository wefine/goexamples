// Example program with Interface, Composition and Method Overriding
package main

import (
	"time"
)

func main() {
	alex := Admin{
		Person{
			"Alex",
			"John",
			time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC),
			"alex@email.com",
			"New York"},
		[]string{"Manage Team", "Manage Tasks"},
	}
	shiju := Member{
		Person{
			"Shiju",
			"Varghese",
			time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			"shiju@email.com",
			"Kochi"},
		[]string{"Go", "Docker", "Kubernetes"},
	}
	chris := Member{
		Person{
			"Chris",
			"Martin",
			time.Date(1978, time.March, 15, 0, 0, 0, 0, time.UTC),
			"chris@email.com",
			"Santa Clara"},
		[]string{"Go", "Docker"},
	}
	team := Team{
		"Go",
		"Golang CoE",
		[]User{alex, shiju, chris},
	}
	//get details of Team
	team.GetTeamDetails()
}
