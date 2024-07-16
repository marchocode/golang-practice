package main

import "fmt"

func main() {

	// 新建一个数组
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	// 从数组生成一个切片
	// [起始位置(默认值0):结束位置(默认值len(array))]
	sli := array[2:5]

	// 切片的容量=len(array)-起始位置
	fmt.Printf("sli =%v length=%d cap=%d\n", sli, len(sli), cap(sli))

	// 声明一个新切片,切片是引用类型，零值为 nil
	var strList []string

	if strList == nil {
		fmt.Println("strList is empty.")
	}

	strList = append(strList, "hello")
	strList = append(strList, "world")

	fmt.Println(strList)

	// length = capacity
	newList := make([]int, 10)
	fmt.Printf("newList len=%d cap=%d \n", len(newList), cap(newList))

	// length=10 capacity=20
	newList = make([]int, 10, 20)
	fmt.Printf("newList len=%d cap=%d \n", len(newList), cap(newList))

	// 测试扩容
	bigSlice := make([]int, 0)

	// 扩容的核心在于 256以内的，每次扩容都翻倍
	// 大于256之后，按照 newcap += (newcap + 3*threshold) >> 2 进行计算容量
	// 但这里有个内存对齐，可能手动算出来的 和实际的容量会有所差异。
	for i := 0; i < 1026; i++ {
		fmt.Printf("bigSlice len=%d cap =%d\n", len(bigSlice), cap(bigSlice))
		bigSlice = append(bigSlice, i)
	}

	next := nextslicecap(513, 512)
	fmt.Println(next)
}

func nextslicecap(newLen, oldCap int) int {

	newcap := oldCap
	doublecap := newcap + newcap
	if newLen > doublecap {
		return newLen
	}

	const threshold = 256
	if oldCap < threshold {
		return doublecap
	}

	for {
		// Transition from growing 2x for small slices
		// to growing 1.25x for large slices. This formula
		// gives a smooth-ish transition between the two.
		newcap += (newcap + 3*threshold) >> 2

		// We need to check `newcap >= newLen` and whether `newcap` overflowed.
		// newLen is guaranteed to be larger than zero, hence
		// when newcap overflows then `uint(newcap) > uint(newLen)`.
		// This allows to check for both with the same comparison.
		if uint(newcap) >= uint(newLen) {
			break
		}
	}

	// Set newcap to the requested cap when
	// the newcap calculation overflowed.
	if newcap <= 0 {
		return newLen
	}
	return newcap
}
