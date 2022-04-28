package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Two-phase Termination Pattern
// 分兩個階段關閉 goroutine，第一階段先結束 goroutine 的程式邏輯，第二階段再結束 goroutine 本身

// 此模式的核心目標如以下三點：
// 安全的終止 goroutine（安全性）
// 必定會關閉 goroutine（生命性）
// 發出關閉請求後盡快運行邏輯關閉的處理（響應性）

// signal.NotifyContext 作法 (Go 1.16+)
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("bye bye")
				os.Exit(1)
			default:
				log.Println("do something")
				time.Sleep(time.Second)
			}
		}
	}()

	<-ctx.Done()

	fmt.Println("shutting down gracefully, press Ctrl+C again to force")
}

// signal.Notify 作法
func example2() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("bye bye")
				os.Exit(1)
			default:
				log.Println("do something")
				time.Sleep(time.Second)
			}
		}
	}()
	// do something
}
