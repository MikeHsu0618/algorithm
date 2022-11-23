package main

import "fmt"

func main() {
	names := [4]string{"小明1", "小明2", "小明3", "小明4"}
	var name = ""

	fmt.Println("請輸入要查找的人名")
	fmt.Scanln(&name)

	// 順序查找： 第一種方式 -> 直接用 i 為基準點
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			fmt.Printf("找到拉 %v %v \n", names[i], i)
			break
		}
		if i == len(names)-1 {
			fmt.Println("真的找不到阿哥")
		}
	}

	// 順序查找： 第二種方式 -> 通過 index 判斷是否找到
	index := -1
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到拉 %v %v \n", names[index], index)
	} else {
		fmt.Println("真的找不到阿哥")
	}
}
