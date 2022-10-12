package main

import (
	"encoding/json"
	"fmt"
)

//"omitempty" tag makes sense only for marshaling from struct to JSON.
//It skips empty values so they won't be in JSON.
//It doesn't affect unmarshaling in any way.
//Use pointers if you want to detect whether the field is specified in JSON or not.
//If the field is not specified, the pointer value will be nil.
func main() {
	test1()
	test2()
}

type Test1 struct {
	String  string `json:"string,omitempty"`
	Integer int    `json:"integer,omitempty"`
}

// omitempty 並不影響 unmarshal 後的值，只影響 marshal 後的 Json 字串不輸出零值

// 但當使用 omitempty 的欄位的值為零值時，也會一併被省略掉
func test1() {
	q := []byte(`{"string":"", "integer": 0}`)
	var qq Test1
	if err := json.Unmarshal(q, &qq); err != nil {
		panic(err)
	}
	fmt.Printf("%#v", qq)

	//fmt.Printf("%#v", *qq.Integer)
	queryStr, err := json.Marshal(qq)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(queryStr)) // {} 輸出為零值
}

type Test2 struct {
	String  *string `json:"string,omitempty"`
	Integer *int    `json:"integer,omitempty"`
}

// 這邊我們可以將 filed 的型別改為指針，使其預設零值為 nil 以防止 omitempty 將其省略
func test2() {
	q := []byte(`{"string":"", "integer": 0}`)
	var qq Test2
	if err := json.Unmarshal(q, &qq); err != nil {
		panic(err)
	}
	fmt.Printf("%#v %#v", *qq.Integer, *qq.String)

	queryStr, err := json.Marshal(qq)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(queryStr)) // {"string":"","integer":0}
}
