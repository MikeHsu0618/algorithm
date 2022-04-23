package main

import (
	"encoding/json"
	"fmt"
)

type user1 struct {
	ID      string `json:"id"`
	Name    string
	Gender  string
	Age     int
	Married bool
	test    string
}

func main() {
	originalUser := user1{
		ID:      "user1",
		Name:    "Mike",
		Gender:  "male",
		Age:     26,
		Married: false,
		test:    "test",
	}

	// 出來的值會是一個 []byte 的格式，所以還要在特地轉成 string
	b, err := json.Marshal(originalUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))

	newUser := user1{}
	err = json.Unmarshal(b, &newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser)
}
