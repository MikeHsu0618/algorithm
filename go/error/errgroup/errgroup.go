package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

// 使用 errorgroup 可以同時堵塞協程並傳入一個 func() error 的方法接收回傳值
func main() {
	// 1. errgroup 底層呼叫 waitGroup 可以堵塞線程並提取錯誤
	eg := errgroup.Group{}
	eg.Go(func() error {
		return fmt.Errorf("error")
	})
	eg.Go(func() error {
		time.Sleep(3 * time.Second)
		fmt.Println("success")
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Printf("get error: %v", err)
	}

	// 2. 第一種方法無法在錯誤發生時，即時停止所有線程，
	// 所以 errgroup 提供了 errgroup.WithContext 方法，使我們傳入 context 並呼叫 cancel
	eg1, ctx := errgroup.WithContext(context.Background())
	// 在這邊我們使其中一個協程出錯並預期第二個協程不執行直接退出
	eg1.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("job cancelled")
			return nil
		default:
			return fmt.Errorf("error happened")
		}
	})
	eg1.Go(func() error {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("job cancelled")
			return nil
		default:
			fmt.Println("success")
			return nil
		}
	})
	if err := eg1.Wait(); err != nil {
		fmt.Println("there were some errors")
	}
}
