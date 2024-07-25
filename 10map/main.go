package main

import "fmt"

func main() {

	m := make(map[string]int, 2)

	m["java"] = 0
	m["goland"] = 1
	m["python"] = 3

	if _, ok := m["goland"]; ok {
		fmt.Println("map contain goland")
	}

	fmt.Println(m)

	delete(m, "java")
	fmt.Println(m)

}
