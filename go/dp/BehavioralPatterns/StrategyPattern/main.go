package main

import "fmt"

// Strategy Pattern
// 設計相同介面但不同實作的物件，再由使用端以此介面去選擇要使用何種物件

// 情境：假設 PS5 會使用不同廠商的顯示晶片，需要有一個方法讓遊戲以不同的顯示晶片運行。

// PS5{}.gpu依賴GPUinterface，A 廠商的AGPU{}與 B 廠商的BGPU{}都是以此GPUinterface 實作，
// 我們在CreatePS5()的時候選定要使用哪個廠商的 GPU，由使用端去選擇要使用何種策略來創造 PS5，選擇完策略後，
// 就可以使用此策略物件來玩遊戲PS5{}.PlayGame()。

type GPU interface {
	Draw()
}

type AGPU struct{}

func (a AGPU) Draw() {
	fmt.Println("draw!")
}

type BGPU struct{}

func (b BGPU) Draw() {
	fmt.Println("draw!")
}

type CGPU struct{}

func (b CGPU) Draw() {
	fmt.Println("draw!")
}

type PS5 struct {
	gpu GPU
}

func CreatePS5(gpu GPU) PS5 {
	ps5 := PS5{
		gpu: gpu,
	}
	return ps5
}

func (p PS5) PlayGame() {
	p.gpu.Draw()
	fmt.Println("play game!")
}

func main() {
	// 在這裡選擇使用的 GPU 後產生不同的 PS5
	gpu := AGPU{} // BGPU{} or CGPU{}
	ps5 := CreatePS5(&gpu)
	ps5.PlayGame()
}
