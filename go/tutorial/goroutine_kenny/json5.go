package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	context1 := `{"id":"user1","name":"Mike","gender":"Male","age":0,"married":false,"Cat":{"name":"any","gender":"Female"},"Dogs":[{"name":"jay","gender":"Male"}],"Hobbies":{"sports":"running, basketball"}}`

	// 如果不確定 json 的格式的話，使用 interface 會幫你轉成 Map[string]interface 的型態
	var i interface{}
	_ = json.Unmarshal([]byte(context1), &i)
	fmt.Println(i)
	m := i.(map[string]interface{})

	// bool for json booleans,
	// float64 for json numbers,
	// string for json strings, and
	// nil for json null.
	for k, v := range m {
		switch v.(type) {
		case float64:
			fmt.Printf("key=%s float value=%v \n", k, v)
		case bool:
			fmt.Printf("key=%s bool value=%v \n", k, v)
		case string:
			fmt.Printf("key=%s string value=%v \n", k, v)
		case nil:
			fmt.Printf("key=%s nil value=%v \n", k, v)
		case []interface{}:
			fmt.Printf("key=%s slice value=%v \n", k, v)
		}
	}

	var i2 interface{}
	content2 := []byte(`["a", 1,false]`)
	_ = json.Unmarshal(content2, &i2)
	fmt.Println(i2)
}
