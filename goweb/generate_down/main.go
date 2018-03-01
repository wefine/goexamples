package main

import (
    "bytes"
    "io"
    "log"
    "math"
    "math/rand"
    "net/http"
    "time"
    "os"
    "fmt"
)

func main() {
    log.SetFlags(log.Lshortfile)
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        modtime := time.Now()
        //content := randomContent(modtime.UnixNano(), 1024)

        // ServeContent uses the name for mime detection
        const name = "random.txt"

        // tell the browser the returned content should be downloaded
        w.Header().Add("Content-Disposition", "Attachment; filename=portal-exports.zip")

        http.ServeContent(w, req, name, modtime, fromZipFile())
    })

    log.Fatal(http.ListenAndServe(":9090", nil))
}

func fromZipFile() io.ReadSeeker {
    file, err := os.Open("c:/mytest.zip")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer file.Close()

    fileInfo, _ := file.Stat()
    var size int64 = fileInfo.Size()

    buffer := make([]byte, size)

    // read file content to buffer
    file.Read(buffer)

    return bytes.NewReader(buffer) // converted to io.ReadSeeker type
}

func randomContent(seed int64, length int) io.ReadSeeker {
    r := rand.New(rand.NewSource(seed))

    content := make([]byte, length, length)
    for i := range content {
        b := byte(r.Intn(math.MaxUint8))

        b = b%('~'-' ') + ' ' // make it a visible character

        content[i] = b
    }

    return bytes.NewReader(content)
}
