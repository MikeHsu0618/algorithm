package main

import "fmt"

func main() {
	// 基本用法
	// var 變數名稱 map[keyType]valueType
	// 聲明範例
	// var a map[string]string
	// var a map[string]map[string]string
	// 注意： 聲明 map 是不會分配內存的，初始化需要 make，分配內存後才能婦職ㄏ和使用

	// 賦值初始化
	var a map[string]string
	a = make(map[string]string, 10)
	a["你好"] = "你好"
	fmt.Println(a)

	b := make(map[string]string, 10)
	b["我是逼"] = "BB"
	fmt.Println(b)

	var c map[string]string = map[string]string{
		"i am C": "CCCCC",
	}
	fmt.Println(c)

	// 案例
	// 存放 3 個學生資料，每個學生有 name 和 age 訊息
	studentMap := make(map[string]map[string]string)

	studentMap["student1"] = map[string]string{
		"name": "John",
		"age":  "18",
	}
	studentMap["student2"] = map[string]string{
		"name": "Tom",
		"age":  "18",
	}
	studentMap["student3"] = map[string]string{
		"name": "Mike",
		"age":  "18",
	}
	fmt.Println(studentMap)

	// map 查找
	val, ok := studentMap["student1"]
	if ok {
		fmt.Println(val, ok)
	} else {
		fmt.Println(val, ok)
	}

	// map 遍歷只能使用 for range
	for key, val := range studentMap {
		fmt.Println(key, val)
		for k, v := range val {
			fmt.Println(k, v)
		}
	}
}
