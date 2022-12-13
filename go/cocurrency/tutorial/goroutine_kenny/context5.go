package main

import (
	"context"
	"fmt"
	"time"
)

// 使用 context.WithDeadline 練習
func main() {
	// 自動三秒過期
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	go task_1(ctx)
	time.Sleep(5 * time.Second)
}

func task_1(ctx context.Context) {
	fmt.Println("task work ...")
	go subTask_1("one", ctx)
	go subTask_1("two", ctx)
	go subTask_1("three", ctx)
}

func subTask_1(name string, ctx context.Context) {
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
