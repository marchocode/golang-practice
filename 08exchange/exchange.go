package exchange

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
