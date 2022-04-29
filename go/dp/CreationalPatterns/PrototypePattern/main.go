package main

import (
	"fmt"
	"time"
)

// Prototype Pattern
// 將物件以拷貝的方式建立

// 如果某類物件在初始化時，會消耗大量的成本，例如：「對 DB 做 query、讀取 json file、讀取使用者滑鼠的軌跡」，
// 那在建立完此物件後，又須建立此此類相似的物件時，我們會希望用拷貝(copy)的，以避免初始化的成本。
// 而拷貝這件事，又有分深拷貝(deep copy) 與 淺拷貝(shallow copy)，如果是有 javascript 經驗的開發者一定很熟悉這兩個詞，

// 在 golang 他們主要是：
// 深拷貝(deep copy)：當拷貝時是拷貝數據本身，所以對拷貝的數據修改，不會影響到原本的數據，
// Int、Float、String、Bool、Struct、Array 都採用這類複製

// 淺拷貝(shallow copy)：當拷貝時是拷貝數據指標，所以對拷貝的數據修改，會影響到原本的數據，Slice、Map 都採用這類複製
// 使用淺拷貝通常是為了更節省記憶體，但如果複製出來的數據有要完全獨立的需求得採用深拷貝，

// 小節一下 Prototype Pattern 的優缺點：
// 優點：
// 不需要初始化的成本
// 缺點：
// 需要注意深淺拷貝的差異，不然有可能會修改到原始的數據

type PS5 struct {
	Version string
	CPU     []string
	GPU     map[string]string
}

func (_ PS5) Modeling() {
	time.Sleep(time.Second) // 模擬耗時
}

func (_ PS5) LoadTest() {
	time.Sleep(time.Second) // 模擬耗時
}

func (_ PS5) Analysis() {
	time.Sleep(time.Second) // 模擬耗時
}

func (p *PS5) Clone() *PS5 {
	// 將基本物件拷貝
	ps5 := PS5{}
	ps5 = *p

	// 由於Slice是淺拷貝，所以必須透過copy方法來把深層元素拷貝
	cpu := make([]string, len(p.CPU))
	copy(cpu, p.CPU)
	ps5.CPU = cpu

	// 由於Map是淺拷貝，所以需把元素一個一個拿出來拷貝
	gpu := make(map[string]string)
	for k, v := range p.GPU {
		gpu[k] = v
	}
	ps5.GPU = gpu

	return &ps5
}

func CreatePrototypePS5() *PS5 {
	prototypePS5 := PS5{
		Version: "Prototype",
		CPU:     []string{"原型CPU"},
		GPU:     make(map[string]string),
	}
	prototypePS5.GPU["GPU"] = "原型GPU"
	prototypePS5.Modeling()
	prototypePS5.LoadTest()
	prototypePS5.Analysis()
	return &prototypePS5
}

func main() {
	prototypePS5 := CreatePrototypePS5()
	ps5 := prototypePS5.Clone()
	ps5.Version = "Version-1"
	ps5.CPU[0] = "量產CPU"
	ps5.GPU["GPU"] = "量產GPU2"
	fmt.Println(prototypePS5)
	fmt.Println(ps5)
}
