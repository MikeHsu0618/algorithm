package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

type MyStruct struct {
	name  string
	value string
}

func main() {
	// 預設過期時間 5 分鐘，每十分鐘檢查是否有過期的 Key
	c := cache.New(5*time.Minute, 10*time.Minute)

	msMap := make(map[string]MyStruct)
	msMap["1"] = MyStruct{
		name:  "123",
		value: "123",
	}
	log.Println(msMap["1"])

	// 存指針可提升效能
	c.Set("map", &msMap, cache.DefaultExpiration)
	if v, found := c.Get("map"); found {
		fmt.Printf("found: %v value: %v \n", found, v)
		value := v.(*map[string]MyStruct)
		fmt.Printf("found: %v value: %v \n", found, value)

	}

}
