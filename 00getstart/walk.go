package main

import (
	"fmt"
	"os"
	"regexp"
	"syscall"
)

var r, _ = regexp.Compile(".*[\\|/]\\.\\w*")

func VisitFile(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}

	if fi.IsDir() {
		return nil // not a file.  ignore.
	}
	fmt.Println(fp)
	return nil
}

func isHidden(path string) (bool, error) {
	p, e := syscall.UTF16PtrFromString(path)
	if e != nil {
		return false, e
	}
	attrs, e := syscall.GetFileAttributes(p)
	if e != nil {
		return false, e
	}
	return attrs&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}

func main() {
	fmt.Println("aa")
}
