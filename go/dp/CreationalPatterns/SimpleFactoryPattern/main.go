package main

import "fmt"

// Simple Factory Pattern
// 由一個工廠將製作產品的細節隱藏，讓使用者不需要不需要知道細節也能獲得產品

// 在製作一個物件的時候，物件常常還需要有其他後製的處理，這些處理使用者不需要知道，
// 使用者只需要獲得此產品即可，所以需要將後製處理作封裝

// 要生產 PS5 主機光碟版與數位版給使用者，但使用者不需要知道「如何生產 CPU、顯示晶片與加裝光碟機」，使用者只需要獲得此產品就行。

type PS5 interface {
	PlayGame()
}

type PS5WithCD struct{}

func (p PS5WithCD) PlayGame() {
	fmt.Println("loading cd...play!")
}
func (p PS5WithCD) AddCDMachine() {
	fmt.Println("adding cd machine...done!")
}
func (p PS5WithCD) AddCPU() {
	fmt.Println("adding cpu...done!")
}
func (p PS5WithCD) AddGPU() {
	fmt.Println("adding gpu...done!")
}

type PS5WithDigital struct{}

func (p PS5WithDigital) PlayGame() {
	fmt.Println("loading digital file...play!")
}
func (p PS5WithDigital) AddCPU() {
	fmt.Println("adding cpu...done!")
}
func (p PS5WithDigital) AddGPU() {
	fmt.Println("adding gpu...done!")
}

// 這裡返回了 PS5 interface 代表返回的物件必須先實現該抽象方法
func CreatePS5(style string) PS5 {
	switch style {
	case "PS5WithCD":
		ps5 := &PS5WithCD{}
		ps5.AddCDMachine()
		ps5.AddCPU()
		ps5.AddGPU()
		return ps5
	case "PS5WithDigital":
		ps5 := &PS5WithDigital{}
		ps5.AddCPU()
		ps5.AddGPU()
		return &PS5WithDigital{}
	}
	return nil
}

func main() {
	// 重點是讓使用者不知道裡面幹了些啥就可得到一台 PS5
	ps5 := CreatePS5("PS5WithCD")
	ps5.PlayGame()
}
