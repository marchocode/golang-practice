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

// Student 类型
type Student struct{ Name string }

// StudentInt 接口类型
type StudentInt interface{}

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
		fmt.Printf("d is not an Integer value.\n")
	}

	// 两个interfaces 相等的前提是 类型相同，且值相同

	// 两个的类型都是 Student, 但值是指针地址，所以不同
	var stu1, stu2 StudentInt = &Student{Name: "1"}, &Student{Name: "1"}

	// 两个的类型都是 Student, 值是结构体，且结构体的值也相同。
	var stu3, stu4 StudentInt = Student{Name: "1"}, Student{Name: "1"}

	fmt.Println("stu1 == stu2", stu1 == stu2)
	fmt.Println("stu3 == stu4", stu3 == stu4)

	var p *int = nil
	var q StudentInt = nil

	// false
	fmt.Println("p == q", p == q)

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
