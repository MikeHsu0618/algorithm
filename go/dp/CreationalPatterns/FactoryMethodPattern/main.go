package main

import "fmt"

// Factory Method Pattern
// 將複雜的生產邏輯再拆分至特定工廠，由使用者端決定要使用什麼工廠來生產產品

// 優點:
// 符合開閉原則，在新增產品時不必修改 function 實作(對修改封閉)，但可以透過不同工廠來新增(對擴充開放)
// 修改都是擴充的，不是改 function 實作，所以把原本的邏輯改壞的可能性不高
// 如果創建邏輯過於複雜，Simple Factory Pattern 的邏輯會相當複雜，
// 而 Factory Method Pattern 將邏輯拆成小工廠可以避免這個問題

// 缺點:
// 為每新增一個產品都需要新增一個工廠，如果產品眾多，程式碼會有數不盡的工廠，讓系統變得很複雜

type GameMachineFactory interface {
	Create() *PS5
}

type PS5WithCDFactory struct{}

func (f PS5WithCDFactory) Create() PS5 {
	ps5 := new(PS5WithCD)
	ps5.AddGPU()
	ps5.AddCDMachine()
	ps5.AddCPU()

	return ps5
}

type PS5WithDigitalFactory struct{}

func (f PS5WithDigitalFactory) Create() PS5 {
	ps5 := new(PS5WithDigital)
	ps5.AddGPU()
	ps5.AddCPU()
	return ps5
}

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

func main() {
	ps5Factory := new(PS5WithCDFactory)
	ps5 := ps5Factory.Create()
	fmt.Println(ps5)
}
