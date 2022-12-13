package main

import "fmt"

type user struct {
	email  string
	name   string
	gender string
}

// 試著將 struct 放進 channel 中
func main() {
	strChan := make(chan string, 3)
	strChan <- "hello world"
	u := user{
		email:  "test@gmail.com",
		name:   "test",
		gender: "male",
	}
	userChan := make(chan user, 1)
	userChan <- u
	fmt.Println(<-userChan)

	// 事實上 channel 預設就是傳遞指針
	// 所以傳入 func 中就不需要 *chan user
	fmt.Println(userChan)
}

type dog struct {
	name, gender string
}

func example2() {
	allChan := make(chan interface{}, 3)
	allChan <- dog{
		name:   "jack",
		gender: "male",
	}
	allChan <- 1
	allChan <- "hello world"
	d := <-allChan
	fmt.Println(d)
}
