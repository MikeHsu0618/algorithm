package main

import "fmt"

// Facade Pattern
// 實作不依賴多個類別(struct)，而是依賴介面(interface)，並把這些類別實作在此介面

// 關鍵是依賴 interface 不依賴 struct，由此可以避免類別替換時，需要修改頂層邏輯

type CPUFacade interface {
	Work()
}

type StructA struct{}

func (s StructA) DoAction() {}

type StructB struct{}

func (s StructB) DoAction() {}

type CPU struct{}

func (c CPU) Work() {
	strcutA := StructA{}
	strcutB := StructB{}
	strcutA.DoAction()
	strcutB.DoAction()
}

type PS5 struct {
	cpu CPUFacade
}

func (p PS5) Start() {
	p.cpu.Work()
	fmt.Println("start ps5...done!")
}

func main() {
	ps5 := PS5{cpu: CPU{}}
	ps5.Start()
}
