package main

import "fmt"

func main() {
	// slice of map
	// 要求：使用一個 map 來紀錄 Monster 訊息的 name 跟 age , 並且可以動態增加 -> map slice
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "Coak"
		monsters[0]["age"] = "400"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "QQQ"
		monsters[1]["age"] = "400"
	}

	// 當超過 make 預設長度（2）時，將會報錯
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string, 2)
	//	monsters[2]["name"] = "QQQ"
	//	monsters[2]["age"] = "400"
	//}

	// 所以我們需要用 append 來動態增長 monsters
	// 1. 定義 monster 資訊
	newMonster := map[string]string{
		"name": "new Monster",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)
}
