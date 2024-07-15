package main

import "fmt"

func main() {

	m := make(map[string]int, 2)

	m["java"] = 0
	m["goland"] = 1
	m["python"] = 3

	fmt.Println(m)

	delete(m, "java")
	fmt.Println(m)

}
