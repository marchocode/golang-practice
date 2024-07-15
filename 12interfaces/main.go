package main

import (
	"fmt"
)

type Human interface {
	color() string
	say() string
}

type Chinese struct {
}

func (c *Chinese) color() string {
	return "yellow"
}

func (c *Chinese) say() string {
	return "中文"
}

func main() {

	var h Human = &Chinese{}

	fmt.Println(h.color())
	fmt.Println(h.say())

	judgeType("demo")
	judgeType(1)
	judgeType(.1)

	var d any

	d = 1
	d = "123"

	if v, ok := d.(int); ok {
		fmt.Printf("d is an Integer value. %v\n", v)
	} else {
		fmt.Printf("d is not an Integer value.")
	}

}

func judgeType(n any) {

	// 该语法只能用于 switch
	switch v := n.(type) {
	case string:
		fmt.Printf("n is string val=%v\n", v)
	case int:
		fmt.Printf("n is string val=%v\n", v)
	default:
		fmt.Printf("unknow type, val=%v\n", v)
	}

}
