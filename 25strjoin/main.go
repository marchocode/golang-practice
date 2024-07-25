package strjoin

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

const letter string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateStr(n int) string {

	b := make([]byte, n)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// 通过+进行拼接字符串
func joinByPlus(n int, str string) string {

	re := ""

	for i := 0; i < n; i++ {
		re += str
	}

	return re
}

func joinByFmt(n int, str string) string {

	re := ""

	for i := 0; i < n; i++ {
		re = fmt.Sprintf("%s:%s", re, str)
	}

	return re
}

// strings.Builder
func joinByBuffer(n int, str string) string {

	buffer := new(strings.Builder)
	// 可以使用预分配的方式
	buffer.Grow(len(str) * n)

	for i := 0; i < n; i++ {
		buffer.WriteString(str)
	}

	// unsafe.String(unsafe.SliceData(b.buf), len(b.buf))
	// 底层都是由 []byte 存储，最后直接转换，并没有开辟空间存储这个新生成的数据
	return buffer.String()
}

// bytes.Buffer
func joinByBytesBuffer(n int, str string) string {

	buf := new(bytes.Buffer)

	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}

	// string(b.buf[b.off:])
	// 最后返回前，将新生成的字符串进行了一次转换保存，
	return buf.String()
}

func joinByBytes(n int, str string) string {

	var buf []byte = make([]byte, 0)

	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}

	return string(buf)
}

func joinByStringsJoin(n int, str string) string {

	// first alloc
	buf := make([]string, n)

	for i := 0; i < n; i++ {
		// 这里需要转换 string 到 byte
		buf[i] = str
	}

	// second.
	return strings.Join(buf, "")
}

func joinByPreBytes(n int, str string) string {

	// 尝试预先分配
	var buf []byte = make([]byte, len(str)*n)

	for i := 0; i < n; i++ {
		// 这里需要转换 string 到 byte
		buf = append(buf, str...)
	}

	return string(buf)
}
