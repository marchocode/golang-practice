package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int, 12)
	var wait sync.WaitGroup
	wait.Add(2)

	// write
	go func() {
		defer wait.Done()

		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("send val =", i)
			time.Sleep(300 * time.Microsecond)
		}

		// 关闭前休眠一段时间
		time.Sleep(2 * time.Second)
		fmt.Println("send val finally")

		defer close(ch)

	}()

	// read
	go func() {

		defer wait.Done()

		for {

			i, ok := <-ch
			if !ok {
				fmt.Println("receive val finally")
				break
			}
			fmt.Println("recive val ", i)
		}

	}()

	wait.Wait()
}
