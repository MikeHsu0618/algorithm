package main

import (
	"fmt"
	"time"
)

// Proxy Pattern
// 讓代理物件操作實際物件，讓代理物件處理與業務邏輯無關的事情

// 在操作物件時，有時候我們有加上 log、權限管控等等功能，這跟實際操作物件的業務無關，
// 這時就可以用 Proxy Pattern 來去處理代理這些功能，並且保持介面相同。

// Proxy Pattern 與 Decorator Pattern 相似，兩者主要差異是：
// Proxy Pattern: 是增加新的功能在原物件上
// Decorator Pattern: 是增強原物件的功能

// PS5 在玩開啟遊戲時，需要新增 log，但如果在原本業務邏輯上加上 log 會模糊掉業務的意圖，我們需要有方法解決此問題。

type PS5 interface {
	PlayGame()
}

type PS5Machine struct{}

func NewPS5Machine() *PS5Machine {
	return &PS5Machine{}
}

func (u *PS5Machine) PlayGame() {
	fmt.Println("play game")
}

type PS5MachineProxy struct {
	ps5 *PS5Machine
}

func NewPS5MachineProxy(ps5 *PS5Machine) PS5 {
	return &PS5MachineProxy{
		ps5: ps5,
	}
}

func (p *PS5MachineProxy) PlayGame() {
	start := time.Now()

	// 在這裡執行原本要執行呢 ps5.PlayGame
	p.ps5.PlayGame()

	// 並在下面加入想要的邏輯 ex. log
	// do something
	
	fmt.Printf("play game cost time: %s", time.Since(start))
}

func main() {
	ps5 := NewPS5MachineProxy(NewPS5Machine())
	ps5.PlayGame()
}
