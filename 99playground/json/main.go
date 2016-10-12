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

	newList := new([10]int)[0:5]
	newList = append(newList,7)
	fmt.Println(newList)

	obj := new(Catalog)
	fmt.Println(obj)

	x := []int{10, 20, 30}
	y := make([]int, 2)
	z := []int{}
	copy(y, x)
	copy(z, x)
	fmt.Println(x, y, z)

	os.Exit(0)

}
