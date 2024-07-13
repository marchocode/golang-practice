package main

import "fmt"

func main() {

	var number int = 10
	var numberPoint *int = &number
	var pointValue = *numberPoint

	fmt.Printf("number val=%d, its address -> %p\n", number, &number)
	fmt.Printf("numberPoint val=%v, its address -> %p\n", numberPoint, &numberPoint)
	fmt.Printf("pointValue val=%v, its address -> %p\n", pointValue, &pointValue)

}
