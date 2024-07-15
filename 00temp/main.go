package main

import "fmt"

func main() {

	// map是引用类型，零值为 nil
	var nilMap map[string]int
	fmt.Printf("nilMap is nil = %v \n", nilMap == nil)

	// 初始化一个空的map
	emptyMap := map[string]int{}
	fmt.Printf("emptyMap's size = %d \n", len(emptyMap))

	// 字面量初始化一个非空map
	initMap := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Printf("initMap = %v \n", initMap)

	// 使用 make创建一个空map
	myMap := make(map[string]int)

	// 新增
	myMap["golang"] = 1
	myMap["python"] = 2
	myMap["rust"] = 3

	fmt.Printf("current myMap's size = %v\n", myMap)

	// 查找
	if v, ok := myMap["golang"]; ok {
		fmt.Printf("myMap contain golang = %d\n", v)
	}

	// 删除
	delete(myMap, "rust")
	fmt.Printf("current myMap = %v \n", myMap)

	// for-each
	// 无序，每次遍历可能不同
	for k, v := range myMap {
		fmt.Printf("current myMap k=%s v=%d\n", k, v)
	}

	sliceMap := &SliceMap{innerSlice: make([]string, 0, 10), innerMap: make(map[string]int)}

	sliceMap.Put("a", 3)
	sliceMap.Put("b", 2)
	sliceMap.Put("c", 1)

	sliceMap.ForEach()

}

type SliceMap struct {
	length     int
	innerSlice []string
	innerMap   map[string]int
}

func (s *SliceMap) Put(k string, v int) {
	s.innerSlice = append(s.innerSlice, k)
	s.innerMap[k] = v
	s.length++
}

func (s *SliceMap) ForEach() {

	for i := 0; i < s.length; i++ {
		k := s.innerSlice[i]
		fmt.Printf("k = %s val=%d \n", k, s.innerMap[k])
	}

}
