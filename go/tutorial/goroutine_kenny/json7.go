package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type user6 struct {
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}

// 如果想要客製 struct json 出來的格式，可以額外實現 MarshqlJSON() 的方法來替代！
func (u *user6) MarshalJSON() ([]byte, error) {
	// 客製輸出的時間格式作法
	// hard code 作法
	// createdAT := u.CreatedAt.Format("2006-01-02 15:04:05")
	// return []byte(fmt.Sprintf(`{"name": "%s","gender":"%s","created_at":"%s"}`, u.Name, u.Gender, createdAT)), nil

	type AliasUser user6
	return json.Marshal(&struct {
		CreatedAt string `json:"created_at"`
		*AliasUser
	}{
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		AliasUser: (*AliasUser)(u),
	})
}

func (u *user6) UnmarshalJSON(data []byte) error {
	type AliasUser user6
	au := struct {
		CreatedAt string `json:"created_at"`
		*AliasUser
	}{
		AliasUser: (*AliasUser)(u),
	}

	if err := json.Unmarshal(data, &au); err != nil {
		return err
	}

	newCreatedAt, err := time.ParseInLocation("2006-01-02 15:04:05", au.CreatedAt, time.Local)
	if err != nil {
		return err
	}

	u.CreatedAt = newCreatedAt

	return nil
}

func main() {
	u := user6{
		Name:      "Mike",
		Gender:    "Male",
		CreatedAt: time.Now(),
	}

	b, _ := json.Marshal(&u)
	fmt.Println(string(b))

	var u2 user6
	_ = json.Unmarshal(b, &u2)
	fmt.Println(u2)
}
