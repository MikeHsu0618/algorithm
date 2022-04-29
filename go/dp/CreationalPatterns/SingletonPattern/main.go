package main

import "sync"

// Singleton Pattern
// 整個程式運作只會有此一個物件，不會創建第二個重複的物件

type DBInstance struct{}

var dbInstance *DBInstance

func InitDB() {
	dbInstance = &DBInstance{}
}

func GetInstance() *DBInstance {
	return dbInstance
}

func GetInstance2() *DBInstance {
	if dbInstance == nil {
		dbInstance = &DBInstance{}
	}
	return dbInstance
}

var (
	once = &sync.Once{}
)

func GetInstance3() *DBInstance {
	once.Do(func() {
		dbInstance = &DBInstance{}
	})
	return dbInstance
}

func main() {
	// 通常作法為在初始入口就實例唯一一次連線，即可重複使用同一個連線
	InitDB()
	GetInstance()

	// Lazy load 等到要使用時才去取得實例
	// 優點：可以在真的需要用時才創建，以節省記憶體
	// 缺點：第一次創建時，會花比較多時間
	GetInstance2()
	// 使用 sync.Once 來實現 Lazy load，可以確保 function 裡的物件永遠只執行一次，以達到 Singleton 的效果
	GetInstance3()
}
