package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//  context.WithCancel 範例2 -> 使用 cancel 跟 done 實現阻塞！
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go task(ctx)

	// 最簡單作法，但無法確定什麼時候要停止
	// time.Sleep(3 * time.Second)
	// fmt.Println("cancel context...")
	// cancel()
	// time.Sleep(3 * time.Second)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGINT)

	// 範例一
	// for 這裡不用用 for {} 因為他會被阻塞等待訊號
	// select {
	// case <-shutdown:
	// 	fmt.Println("shutdown")
	// 	cancel()
	// 	time.Sleep(2 * time.Second)
	// 	return
	// }

	// 範例二
	<-shutdown
	fmt.Println("shutdown")
	cancel()
	time.Sleep(2 * time.Second)
}

func task(ctx context.Context) {
	fmt.Println("task work...")
	go subTask("one", ctx)
	go subTask("two", ctx)
	go subTask("three", ctx)
}

func subTask(name string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("subTask %s stop work by ctx cancel, err: %v \n", name, ctx.Err())
			return
		default:
			fmt.Printf("subTask %s work \n", name)
			time.Sleep(time.Second)
		}
	}
}
