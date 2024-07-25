package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	// a -> U+0061
	a := 'a'
	// a -> 1100001
	fmt.Printf("a = %b, type =%s \n", a, reflect.TypeOf(a).Kind())

	// a -> 单字节
	// [01100001 0 0 0 0 0 0 0 0 0 ]
	printFile("a.txt")

	// 严 -> U+4E25
	// 严 -> 100111000100101
	fmt.Printf("严 = %b \n", '严')

	// 编码后三个字节
	// [11100100 10111000 10100101 0 0 0 0 0 0 0 ]
	// 连续3个1 说明编码后的字符占三个字节
	printFile("b.txt")

}

func printFile(name string) {

	file1, err := os.Open(name)
	bytes1 := make([]byte, 10)

	if err != nil {
		fmt.Println("input error")
	}

	file1.Read(bytes1)
	byteToStr(bytes1)
}

func byteToStr(bytes []byte) {
	fmt.Print("[")
	for _, val := range bytes {
		fmt.Printf("%b ", val)
	}
	fmt.Print("]\n")
}
