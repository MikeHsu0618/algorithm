package main

// wait group 可以很安全的處理 gorutine 的資源爭搶問題
// 常用函數有 Add Done Wait
func main() {
	//  wait group 的操作都是使用指針操作

	// 第一個案例可以宣告為值引用，並且還能正常的 call function
	// 是因為 golang 的語法糖可以將值傳遞轉成指針傳遞
	// note 通常 call struct 會用指針傳遞而不會使用值傳遞，
	// 第一個原因耗資源，第二個原因可以在其他地方共同操作同個 struct

	// var wg1 sync.WaitGroup
	// wg2 := new(sync.WaitGroup)
	// wg3 := &sync.WaitGroup{}
	// var wg4 = &sync.WaitGroup{}
}
