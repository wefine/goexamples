package main

import (
    "os"
    "path/filepath"
    "strings"
    "io/ioutil"
    "log"
    "unicode/utf8"
)

func main() {
    dirPath := os.Args[1]

    // walk all files in directory
    filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() && strings.HasSuffix(info.Name(), ".java") {

            buf, err := ioutil.ReadFile(path)
            if err != nil {
                log.Fatal(err)
            }

            size := 0
            for start := 0; start < len(buf); start += size {
                var r rune
                if r, size = utf8.DecodeRune(buf[start:]); r == utf8.RuneError {
                    log.Println("invalid utf8 encoding:", path)
                    break
                }
            }
        }
        return nil
    })
}
