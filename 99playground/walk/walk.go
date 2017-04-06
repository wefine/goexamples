package main

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path, _ = os.Getwd()
	}

	read_link(path)

	fmt.Println("done!")
}

func read_dir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if file.Mode() == os.ModeSymlink {
			println("link")
		}

		fileInfo, err := os.Lstat(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Link info: %+v", fileInfo)
	}
}

func read_link(path string) {
	exit_code := 0
	defer os.Exit(exit_code)

	ln, err := filepath.EvalSymlinks(path)
	if err != nil {
		fmt.Println("[ERR]", err)
		exit_code = 1
		return
	}

	fmt.Println("[FOUND]", ln)
}

func walk() {
	dirPath := os.Args[1]
	// walk all files in directory
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			println(info.Name())
			if info.Mode() == os.ModeSymlink {
				println("link")
			}
		}
		return nil
	})
}

func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}
