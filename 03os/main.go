package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	args := os.Args[1:]

	output, split := "", ""

	st1 := time.Now()
	for i := 0; i < len(args); i++ {
		output += split + args[i]
		split = " "
	}

	fmt.Println("first is over. ", time.Now().Sub(st1))

	fmt.Println(output)
	fmt.Println(args)

	st2 := time.Now()
	fmt.Println(strings.Join(args, " "))
	fmt.Println("second is over. ", time.Now().Sub(st2))

	for index, value := range os.Args[1:] {
		fmt.Println("index", index, "value", value)
	}

	fmt.Println(time.Now())

}
