package main
import "flag"
import "fmt"

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
}

// go run c1-command-args-parser.go -i in.dat -o out.dat -a bubble
// infile = in.dat outfile = out.dat algorithm = bubble
