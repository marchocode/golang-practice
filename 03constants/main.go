package main

import "fmt"

const (
	SUNDAY  = iota //0
	MONDAY         //1
	TUESDAY        //2
)

func main() {

	const PI = 3.14
	fmt.Printf("PI type %T value=%v \n", PI, PI)

	fmt.Printf("Sunday = %v \n", SUNDAY)
	fmt.Printf("MONDAY = %v \n", MONDAY)
	fmt.Printf("TUESDAY = %v \n", TUESDAY)
}
