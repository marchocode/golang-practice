package main

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var strs []string = []string{""}
	item := longestCommonPrefix(strs)
	fmt.Println(item)
}
