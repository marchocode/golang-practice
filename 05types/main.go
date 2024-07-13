package main

import "fmt"

type Integer = int
type Name string

func main() {

	var a Integer = 10
	fmt.Printf("type a = %T \n", a)

	var b Name = "Marcho"
	fmt.Printf("type b = %T \n", b)

}
