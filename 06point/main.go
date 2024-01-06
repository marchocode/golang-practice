package main

import "fmt"

func main() {

	var a int = 1
	var b *int = &a

	fmt.Println(*b)

	var c = new(int)
	fmt.Println(*c)

}
