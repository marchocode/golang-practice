package strjoin

import (
	"testing"
)

func benchmark(b *testing.B, f func(int, string) string) {
	var str = generateStr(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

// fmt.Sprintf("%s:%s", re, str)
// 这种性能最差，一般用于格式化字符串，不用于拼接
func BenchmarkJoinByFmt(b *testing.B) {
	benchmark(b, joinByFmt)
}

// 通过+拼接，性能差
func BenchmarkJoinByPlus(b *testing.B) {
	benchmark(b, joinByPlus)
}

// strings.Builder
// buffer.Grow(len(str) * n)
// 可以通过预分配的方式，增加其性能
func BenchmarkJoinByBuffer(b *testing.B) {
	benchmark(b, joinByBuffer)
}

// bytes.Buffer
func BenchmarkJoinByBytesBuffer(b *testing.B) {
	benchmark(b, joinByBytesBuffer)
}

// 
func BenchmarkJoinByBytes(b *testing.B) {
	benchmark(b, joinByBytes)
}

func BenchmarkJoinByPreBytes(b *testing.B) {
	benchmark(b, joinByPreBytes)
}

func BenchmarkJoinByStringsJoin(b *testing.B) {
	benchmark(b, joinByStringsJoin)
}
