package main

import "fmt"

func init() {

}

func main() {

	var a uint = 1
	var b uint = 2
	var c uint = 1<<64 - 1

	var d uint = 5
	var e uint = 10

	// uint是不包含负数的，如果减法出现负数，会出现下溢。绕过0后到达最大值
	fmt.Println(a - b)

	fmt.Println("0 = ", 0)
	fmt.Println("-1 = ", c)
	fmt.Println("-2 = ", c-1)
	fmt.Println("-3 = ", c-2)
	fmt.Println("-4 = ", c-3)
	fmt.Println("-5 = ", c-4)

	// 5-10=-5
	fmt.Println((d - e) == (1<<64 - 5))

}
