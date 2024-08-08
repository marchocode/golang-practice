package datastruct

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 9
	re := twoSum(nums, target)
	t.Log(re)
}

func TestTwoSumHash(t *testing.T) {

	nums := []int{3, 2, 4}
	target := 6
	re := twoSumHash(nums, target)
	t.Log(re)
}

func TestSearch(t *testing.T) {

	nums := []int{1, 3, 5, 6}
	target := 5

	re := SearchInsert(nums, target)
	fmt.Println(re)
}

func TestMerge(t *testing.T) {
	nums := [][]int{{1, 4}, {0, 4}}
	li := Merge(nums)
	fmt.Println(li)
}
