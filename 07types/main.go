package main

import "fmt"

type Kilogram int
type Gram int

func (kg Kilogram) String() string {
	return fmt.Sprintf("%d Kg", kg)
}

func (g Gram) String() string {
	return fmt.Sprintf("%d g", g)
}

func GramToKilo(d Gram) Kilogram {
	return Kilogram(d / 1000)
}

func main() {

	var kg = Kilogram(10)

	fmt.Println(kg.String())
	fmt.Println(GramToKilo(1000))
}
