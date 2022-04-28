package main

import (
	"context"
	"fmt"
	"time"
)

// Thread-Specific Storage Pattern
// 為每個 goroutine 擁有自己的儲存空間，供不同的情境識別與使用

// 所以 golang 是採用 context 的做法來達到「讓 goroutine 擁有自己的儲存空間」，
// context 可以利用With等方法，把 context 夾帶某些值，帶到不同的 goroutine，

// context.WithValue(ctx, key, val): 可以將原有 context 新增某 key，並存放著某 value

type RequestID struct{}

func RequestHandler(ctx context.Context) {
	fmt.Printf("request ID is %v\n", ctx.Value(RequestID{}))

	// do something
}

func main() {
	ctx := context.Background()
	go RequestHandler(context.WithValue(ctx, RequestID{}, 1))
	go RequestHandler(context.WithValue(ctx, RequestID{}, 2))
	go RequestHandler(context.WithValue(ctx, RequestID{}, 3))

	time.Sleep(10 * time.Second) // 等待goroutine執行完畢
}
