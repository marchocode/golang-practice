package main

import (
	"example/greeting"
	"fmt"
)

func main() {
	message := greeting.Hello("golang")
	fmt.Print(message)
}
