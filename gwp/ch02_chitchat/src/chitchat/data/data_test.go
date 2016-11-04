package data

import (
	"fmt"
	"testing"
)

// test data
var users = []User{
	{
		Name:     "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
}

func setup() {
	ThreadDeleteAll()
	SessionDeleteAll()
	UserDeleteAll()
}

func Test_createUUID(t *testing.T) {
	uuid := createUUID()
	fmt.Println(uuid)
}
