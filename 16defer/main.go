package main

import (
	"errors"
	"fmt"
)

// defer 在return之前会依次出栈
func demo() int {

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	return 6
}

func hasErr() (int, error) {

	// defer要定义在error之前，否则不生效
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	if true {
		return 0, errors.New("some error")
	}

	defer fmt.Println("4")

	return 6, nil
}

func main() {

	num := demo()
	fmt.Printf("demo() return val=%d \n", num)

	num2, err := hasErr()

	if err != nil {
		fmt.Printf("hasErr() return val=%d \n", num2)
	}

}
