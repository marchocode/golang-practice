package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "desc.")
var sep = flag.String("s", " ", "splite string.")

func main() {

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}

}
