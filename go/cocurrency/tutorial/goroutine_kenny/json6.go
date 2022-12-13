package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	// 執行後可以在 cli 上輸入，如果為 json 格式則自動 Encode 為 Json
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
