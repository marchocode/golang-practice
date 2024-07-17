package main

import "fmt"

func money(m float64) {

	if m > 100 {
		panic("余额不足")
	}

}

func main() {

	defer func() {
		// 通过定义recover函数进行恢复
		if err := recover(); err != nil {
			// 执行恢复逻辑
			fmt.Println("recover process.")
		}

	}()

	// panic
	money(200)

	// does't output.
	fmt.Println("success")

	// 如果需要继续执行函数，通过匿名函数的方式，将panic进行包裹

}
