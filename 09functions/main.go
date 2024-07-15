package main

import "fmt"

// 可以将函数定义为一个类型
type Customer func(item any)

func main() {

	array := []int{1, 2, 3, 4}

	// 函数式编程
	// 将一个匿名函数作为参数进行传递
	forEach(array, func(item any) { fmt.Print(item) })
	fmt.Println()

	// 闭包
	counter := generateCounter()

	fmt.Println(counter()) //1
	fmt.Println(counter()) //2
	fmt.Println(counter()) //3

}

// 通过函数返回一个函数，并且这个返回的函数可以携带其作用域外的一些变量
func generateCounter() func() int {

	num := 0

	return func() int {
		num++
		return num
	}
}

func forEach(arrays []int, cust Customer) {

	for _, val := range arrays {
		cust(val)
	}

}
