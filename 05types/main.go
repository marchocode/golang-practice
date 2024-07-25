package main

import "fmt"

// 别名
type Integer = int
// 新的类型
type Name string

func main() {

	var a Integer = 10
	fmt.Printf("type a = %T \n", a)

	var b Name = "Marcho"
	fmt.Printf("type b = %T \n", b)

}
