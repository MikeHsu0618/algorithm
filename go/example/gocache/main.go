package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type MyStruct struct {
	name  string
	value string
}

func main() {
	// 預設過期時間 5 分鐘，每十分鐘檢查是否有過期的 Key
	c := cache.New(5*time.Minute, 10*time.Minute)

	testMap := make(map[string]string)
	testMap["foo"] = "barrr"

	c.Set("foo", "bar", 5*time.Second)
	c.Set("baz", 42, cache.NoExpiration)
	c.Set("testMap", testMap, cache.NoExpiration)

	v, found := c.Get("foo")
	fmt.Printf("found: %v value: %v \n", found, v)

	v, found = c.Get("baz")
	fmt.Printf("found: %v value: %v \n", found, v)

	v, found = c.Get("testMap")
	fmt.Printf("found: %v value: %v \n", found, v) // found

	time.Sleep(5 * time.Second)

	v, found = c.Get("foo")
	fmt.Printf("found: %v value: %v \n", found, v) // not found

	// 實際用法
	ms := &MyStruct{
		name:  "test",
		value: "value",
	}

	// 存指針可提升效能
	c.Set("mystruct", &ms, cache.DefaultExpiration)
	if v, found := c.Get("mystruct"); found {
		fmt.Printf("found: %v value: %v \n", found, v)
	}
}
