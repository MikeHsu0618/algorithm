package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//  context.WithCancel 範例 -> 必須從頂層傳入喔！
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	luckyNumber := 5
	go guess(ctx, luckyNumber)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
}

func guess(ctx context.Context, luckyNumber int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx cancel,cant guess right number")
			return
		default:
			if num := rand.Intn(5) + 1; num == luckyNumber {
				fmt.Println("guess right luckyNumber:", num)
				return
			} else {
				fmt.Println("guess wrong luckyNumber:", num)
			}
		}
	}
}
