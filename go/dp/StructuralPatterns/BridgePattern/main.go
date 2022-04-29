package main

// Bridge Pattern
// 將抽像與實現分離，讓彼此變化互不影響

// 情境：如果要將兩種 PS5Machine 接上兩種控制器，卻不想要在兩個 PS5Machine 上各新增兩種控制器，
// 可以使用 Bridge Pattern 將控制器抽象出來當作一個介面，使其階層與 PS5Machine 平行，
// 並且由上層的 PS5 來負責管理。

type Controller interface {
	MethodA()
	MethodB()
	MethodC()
	MethodD()
}

type ControllerA struct{}

func (c ControllerA) MethodA() {}
func (c ControllerA) MethodB() {}
func (c ControllerA) MethodC() {}
func (c ControllerA) MethodD() {}

type ControllerB struct{}

func (c ControllerB) MethodA() {}
func (c ControllerB) MethodB() {}
func (c ControllerB) MethodC() {}
func (c ControllerB) MethodD() {}

type PS5 interface {
	Start()
	SetController()
	Play()
}

type PS5Machine struct {
	ps5Controller Controller
}

func (p PS5Machine) Start() {
	p.ps5Controller.MethodA()
	p.ps5Controller.MethodB()
}

func (p *PS5Machine) SetController(controller Controller) {
	p.ps5Controller = controller
}

func (p PS5Machine) Play() {
	p.ps5Controller.MethodC()
	p.ps5Controller.MethodD()
}

func main() {
	ps5 := PS5Machine{}
	ps5.Start()
	ps5.SetController(ControllerA{})
	ps5.Play()
	ps5.SetController(ControllerB{})
	ps5.Play()
}
