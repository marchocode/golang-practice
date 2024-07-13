package main

import (
	"fmt"
	"os"
)

// 多行定义
// 函数外部定义的变量作用域是全局，全局变量必须以var进行声明
var (
	name string
	age  int
)

func main() {

	fmt.Printf("Default name type = %T\n", name)
	fmt.Printf("Default age type = %T\n", age)

	// 局部变量，作用域为该函数内部
	// define a variable.
	// 可同时声明多个变量，未进行赋值操作，默认值为零值。
	var x, y int
	fmt.Printf("x value =%d, y value = %d \n", x, y)

	// 声明并初始化
	var cost float64 = 10.2
	fmt.Printf("we cost %f$ \n", cost)

	// 省略类型，编译器推断类型
	var name = "Marcho"
	fmt.Printf("My name is %s \n", name)

	// 简短的声明赋值操作，但是只能用于函数内部
	address := "something.."
	fmt.Printf("Your address is %s \n", address)

	// 匿名变量
	file, _ := os.Open("go.mod")
	fmt.Printf("opne a file %p\n", file)
}

// 这里的a b 是形参，即这个函数被调用的时候才会分配对应的值
// 函数运行结束后会被销毁
// 在没有调用的情况下，不占用任何空间
// 在函数调用的时候，形参会转变为实参，且作用域是该函数内部
func add(a int, b int) int {
	return a + b
}
