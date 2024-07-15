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

	// 声明一个新切片
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

	// copy
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	// 数组大小不一致，按照较小的那个进行复制
	copy(slice2, slice1)
	fmt.Println(slice2)

	// 复制影响
	const elementCount = 20
	srcData := make([]int, elementCount)

	// 赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	fmt.Println(srcData)

	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	refData := srcData
	fmt.Println(refData)

	// 复制一份
	copyData := make([]int, elementCount)
	copy(copyData, srcData)
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	fmt.Println(copyData)

	// 修改原数据
	srcData[0] = 99
	// [99 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	fmt.Println(srcData)
	// [99 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	fmt.Println(refData)

	// 由于是复制的，不会发生变化
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	fmt.Println(copyData)

	// 修改copyData
	copyData[len(copyData)-1] = 66
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 66]
	fmt.Println(copyData)

	// 不会影响
	fmt.Println(srcData)

	// 尝试在其他函数中改变切片的值
	change(srcData)

	// 已经发生改变，因为切片是一个引用类型。
	fmt.Println(srcData)
	fmt.Printf("src point = %p\n", srcData)
}

func change(s []int) {
	s[1] = 88
}
