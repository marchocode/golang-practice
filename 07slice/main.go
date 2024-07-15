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

	// 扩容测试
	newSlice := make([]int, 0, 1)
	// 1024 以内
	for i := 0; i < 10; i++ {
		fmt.Printf("newSlice len=%d cap =%d val=%v\n", len(newSlice), cap(newSlice), newSlice)
		newSlice = append(newSlice, i)
	}
	// 256以外
	bigSlice := make([]int, 1024, 1024)

	newLen := 1025
	threshold := 256
	size := 1024
	newcap := size

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

	fmt.Printf("resize =%d\n", newcap)

	for i := 0; i < 10; i++ {
		fmt.Printf("bigSlice len=%d cap =%d\n", len(bigSlice), cap(bigSlice))
		bigSlice = append(bigSlice, i)
	}

	// 尝试在其他函数中改变切片的值
	change(srcData)

	// 已经发生改变，因为切片是一个引用类型。
	fmt.Println(srcData)
	fmt.Printf("src point = %p\n", srcData)
}

func change(s []int) {
	s[1] = 88
}
