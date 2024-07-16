package main

import "fmt"

// 变量逃逸
// moved to heap: num
func f1() *int {

	num := 10
	return &num

}

func main() {

	f1()

	// num2在编译过程中，其类型无法确定
	// num2 escapes to heap
	num2 := 10
	fmt.Println("value =", num2)

	// 较小
	// make([]int, 0, 1) does not escape
	sli := make([]int, 0, 1)

	for i := 0; i < cap(sli); i++ {
		sli[i] = i
	}

	// 较大变量，逃逸到栈上
	// make([]int, 0, 10000) escapes to heap

	sli2 := make([]int, 0, 10000)

	for i := 0; i < cap(sli2); i++ {
		sli2[i] = i
	}

	// 逃逸，43行 sliCap 无法确定其大小
	// make([]int, sliCap) escapes to heap
	sliCap := 10
	sli3 := make([]int, sliCap)

	for i := 0; i < cap(sli3); i++ {
		sli3[i] = i
	}

}
