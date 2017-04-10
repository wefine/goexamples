package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"os/exec"
)

func main() {
	var path string

	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		path, _ = os.Getwd()
	}

	findLnInDir(getFullPath(path))
}

func findLnInDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(dir, " is not a directory!")
		return
	}

	var path string
	for _, fi := range files {
		if len(fi.Name()) < 3 {
			// skip . and ..
			continue
		}

		path = getFullPath(dir + "/" + fi.Name())

		if fi.Mode()&os.ModeSymlink != 0 {
			replaceLink(path)
		} else if fi.IsDir() {
			findLnInDir(path)
		}
	}
}

func replaceLink(link string) {
	ln, _ := filepath.EvalSymlinks(link)
	realPath := getFullPath(ln)

	path := getFullPath(link)

	log.Println("realPath=" + realPath +"; path=" + path)
return

	err := copyCmd(realPath, path)
	if err != nil {
		log.Println("Can not handle file: ", path, err)
	}
}

func getFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func copyCmd(src, dst string) (error) {
	_, err := os.Stat(src)
	if err != nil {
		return err
	}

	execCmd("rm -rf " + dst)
	execCmd("cp -rf " + src + " " + dst)

	return nil
}

func execCmd(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Printf("Failed to execute command: %s", cmd)

		return cmd
	}

	return string(out)
}