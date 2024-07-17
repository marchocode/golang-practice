package main

import (
	"fmt"
	"runtime"
	"sync"
)

func task1(wait *sync.WaitGroup) {
	for i := 0; i < 50; i++ {
		fmt.Printf("run task1 num=%d \n", i)
	}

	if wait != nil {
		wait.Done()
	}
}

func task2(wait *sync.WaitGroup) {
	for i := 0; i < 50; i++ {
		fmt.Printf("run task2 num=%d \n", i)
	}
	if wait != nil {
		wait.Done()
	}
}

func main() {

	// cpu核心数量，runtime.GOMAXPROCS() 的默认值
	fmt.Printf("cpu number = %d\n", runtime.NumCPU())

	// 串行
	// task1(nil)
	// task2(nil)
	var group sync.WaitGroup

	// task number.
	group.Add(2)

	// 并发
	go task1(&group)
	go task2(&group)

	group.Wait()
}
