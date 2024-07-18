package main

import (
	"context"
	"fmt"
	"time"
)

func cancelTask(ctx context.Context) {

	for {
		select {

		case <-ctx.Done():
			fmt.Println("cancelTask end.")
			return
		default:
			fmt.Println("cancelTask running.")
			time.Sleep(1 * time.Second)
		}
	}

}

func cancelTaskTest() {

	// root context.
	root := context.Background()

	ctx, cancel := context.WithCancel(root)
	go cancelTask(ctx)

	// main thread.
	time.Sleep(3 * time.Second)
	cancel()

	// done.
	time.Sleep(2 * time.Second)
}

func timeoutTaskTest() {

	// root context.
	root := context.Background()

	// 超时自动取消
	ctx, cancel := context.WithTimeout(root, 3*time.Second)
	defer cancel()

	go cancelTask(ctx)

	// main thread.
	time.Sleep(5 * time.Second)
	fmt.Println("main process end.")
}




func main() {

	// cancelTaskTest()

	timeoutTaskTest()
}
