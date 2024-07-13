package main

import "fmt"

func main() {

	// 初始化一个定长的数组,并且给前三个进行赋值,未初始化的值就是0
	var arr [4]int = [4]int{1, 2, 3}
	fmt.Println(arr)

	// 不给定长度，按照给定的初始化值自动计算
	// type = [6]int
	log := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(log)
	fmt.Printf("log type = %T\n", log)
}
