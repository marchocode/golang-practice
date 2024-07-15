package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// 指定范型
func max[T int | float64 | float32 | int64](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func main() {

	fmt.Printf("min value = %d\n", min(3, 5))
	fmt.Printf("max value = %v\n", max(3, 5))
	fmt.Printf("max value = %v\n", max(3.6, 2.5))

}
