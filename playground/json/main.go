package main

import (
	"fmt"
	"os"
)

type Catalog struct {
	Repositories []string `json:"repositories"`
}

type ImageInfo struct {
	Name string `json:"name"`
	Tags []string `json:"tags"`
}

func main() {
	makeList := make([]Catalog, 5, 10)
	fmt.Println(makeList)

	newList := new([10]Catalog)[0:5]
	fmt.Println(newList)

	obj := new(Catalog)
	fmt.Println(obj)
	os.Exit(0)

}
