package main

import "runtime"

func main() {
	// 取得可用 CPU 核心數
	cpuNum := runtime.NumCPU()
	println("可使用的 CPU 數 = ", cpuNum)

	// 可以自己設置使用多少個 CPU (留一個給其他人用)
	runtime.GOMAXPROCS(cpuNum - 1)

	// go1.8 後, 默認讓程序運行在多個核心上, 可以不用設置了
	// go1.8 前, 還是要設置一下, 可以更高效使用 CPU
}
