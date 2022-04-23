package main

import (
	"encoding/json"
	"fmt"
)

type user4 struct {
	ID      string `json:"-"`                // 不轉成 json 內容
	Name    string `json:"name"`             // 正常改變 json key 值
	Gender  string `json:"gender,omitempty"` // omitempty —> 如為空值則不回傳
	Age     int    `json:"-,"`               // 反骨，就是想把 Key 命名為 "-"
	Married bool   `json:",omitempty"`       // 不該變回傳預設 Key 命名，但空值則不回傳
}

type user5 struct {
	Name   string `json:"name"`
	Gender string `json:"gender,omitempty"`
}

func main() {
	context1 := `{"id":"user1","Name":"Mike","Gender":"male","Age":26,"Married":false}`
	context2 := `{"id":"user1","Name":"Mike","Gender":"male","Age":26,"Married":false}`

	// json.Valid 會幫我們檢查 json string 格式，並幫我們過濾多餘的欄位
	fmt.Println(json.Valid([]byte(context1)))
	fmt.Println(json.Valid([]byte(context2)))

	var u1 user4
	var u2 user5

	_ = json.Unmarshal([]byte(context1), &u1)
	fmt.Println(u1)

	_ = json.Unmarshal([]byte(context2), &u2)
	fmt.Println(u2)
}
