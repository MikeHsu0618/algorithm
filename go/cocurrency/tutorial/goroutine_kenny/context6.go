package main

import (
	"context"
	"fmt"
	"time"
)

// 使用 context.WithValue 範例
func main() {
	ctx := context.WithValue(context.Background(), "username", "Mike")
	go task_2(ctx)
	time.Sleep(3 * time.Second)
}

func task_2(ctx context.Context) {
	fmt.Println("task get value", ctx.Value("username"))
}
