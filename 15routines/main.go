package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(w *sync.WaitGroup) {

	defer w.Done()

	for i := 0; i < 50; i++ {
		fmt.Printf("run task1 number=%d \n", i)
	}

}

func task2(w *sync.WaitGroup) {

	defer w.Done()

	for i := 0; i < 50; i++ {
		fmt.Printf("run task2 number=%d \n", i)
	}

}

// 生产者
func producter(c chan int) {

	for i := 0; i < 10; i++ {
		fmt.Printf("producter number=%d \n", i)
		c <- i
		time.Sleep(time.Microsecond * 200)
	}

	close(c)
}

// 消费者
func consumer(c chan int) {

	// 从channel中读取数据
	for i := range c {
		fmt.Printf("consumer get number=%d \n", i)
		time.Sleep(time.Microsecond * 300)
	}

}

var x int

func addNoLock() {

	for i := 0; i < 5000; i++ {
		x = x + 1
	}

}

// 互斥锁
var lock sync.Mutex

func addLock() {

	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}

}

// 读写锁
var rwLock sync.RWMutex

func download() {

	for i := 0; i <= 100; i++ {
		rwLock.Lock()
		x = x + 1
		rwLock.Unlock()
		time.Sleep(time.Microsecond * 100)
	}

}

func downloadProcess() {

	for {
		rwLock.Lock()
		fmt.Printf("download process = %d %%\n", x)

		if x == 100 {
			rwLock.Unlock()
			break
		}
		rwLock.Unlock()
	}

}

func main() {

	// 并发线程2

	// var group sync.WaitGroup
	// group.Add(2)

	// go task1(&group)
	// go task2(&group)

	// group.Wait()

	// 生产者与消费者的通道通信
	// ch := make(chan int, 5)

	// go producter(ch)
	// go consumer(ch)

	// 不加锁，多个协程访问临界区，会造成竞争问题
	// go addNoLock()
	// go addNoLock()

	// time.Sleep(time.Second * 5)
	// fmt.Printf("x = %d \n", x)

	// 加锁，互斥锁完全读写互斥
	// 互斥锁能够保证同一时间有且只有一个goroutine进入临界区
	// go addLock()
	// go addLock()

	// time.Sleep(time.Second * 5)
	// fmt.Printf("x = %d \n", x)

	// 读写锁
	// 当读多写少的场景，使用读写锁可以提供性能
	go download()
	go downloadProcess()

	time.Sleep(time.Second * 5)
}
