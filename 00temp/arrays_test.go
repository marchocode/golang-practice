package temp

import (
	"fmt"
	"testing"
)

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
