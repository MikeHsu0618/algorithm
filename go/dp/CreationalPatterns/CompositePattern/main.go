package main

import "fmt"

// Composite Pattern
// 將單一與多個物件的使用方式統一給使用者使用

type CPU interface {
	Run()
}

type SingleCPU struct{}

func (SingleCPU) Run() {
	fmt.Println("run cpu")
}

type MultiCPUs struct {
	SubCPUs []CPU
}

func (m *MultiCPUs) Run() {
	for _, cpu := range m.SubCPUs {
		cpu.Run()
	}
}

func (m *MultiCPUs) AddSubCPU(cpu CPU) {
	m.SubCPUs = append(m.SubCPUs, cpu)
}

func PS5Start(cpu CPU) {
	cpu.Run()
}

func main() {
	singleCPU1 := SingleCPU{}
	PS5Start(singleCPU1)

	singleCPU2 := SingleCPU{}
	PS5Start(singleCPU2)

	multiCPUs := MultiCPUs{}
	multiCPUs.AddSubCPU(&singleCPU1)
	multiCPUs.AddSubCPU(&singleCPU2)
	PS5Start(&multiCPUs)
}
