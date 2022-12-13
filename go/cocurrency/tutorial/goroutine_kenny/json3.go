package main

import (
	"encoding/json"
	"fmt"
)

type Cat1 struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type user3 struct {
	ID      string `json:"-"`                // 不轉成 json 內容
	Name    string `json:"name"`             // 正常改變 json key 值
	Gender  string `json:"gender,omitempty"` // omitempty —> 如為空值則不回傳
	Age     int    `json:"-,"`               // 反骨，就是想把 Key 命名為 "-"
	Married bool   `json:",omitempty"`       // 不該變回傳預設 Key 命名，但空值則不回傳
	Cat     *Cat1  // 如果想要設定 nil 的話可以使用指針形式，因為指針預設就是 nil !
}

func main() {
	originUser := user3{
		ID:      "user1",
		Name:    "Mike",
		Gender:  "Male",
		Age:     0,
		Married: false,
		Cat:     nil,
	}
	b, _ := json.MarshalIndent(originUser, "", "    ")

	fmt.Println(string(b))
}
