package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type CustomDateTimeFormat string

func (c CustomDateTimeFormat) rfc3330Format() string {
	return time.RFC3339
}

func (c CustomDateTimeFormat) myFormat() string {
	return "2006-01-02 15:04:05"
}

type CustomDateTime struct {
	time.Time
	CustomDateTimeFormat
}

func (d *CustomDateTime) MarshalJSON() ([]byte, error) {
	dateTimeStr := d.Time.Format(d.CustomDateTimeFormat.myFormat())
	return []byte(fmt.Sprintf(`"%s"`, dateTimeStr)), nil
}

func (d *CustomDateTime) UnmarshalJSON(b []byte) error {
	fmt.Println(b, "test")
	b = bytes.Trim(b, `"`)
	dataTime, _ := time.Parse(d.CustomDateTimeFormat.myFormat(), string(b))
	d.Time = dataTime
	return nil
}

type user7 struct {
	Name      string         `json:"name"`
	Gender    string         `json:"gender"`
	CreatedAt CustomDateTime `json:"created_at"`
}

// 客製出自己的 Marshal Unmarshal 的方法
func main() {
	u := user7{
		Name:   "Mike",
		Gender: "Male",
		CreatedAt: CustomDateTime{
			Time: time.Now(),
		},
	}

	b, _ := json.Marshal(&u)
	fmt.Println(string(b))

	var u2 user7
	_ = json.Unmarshal(b, &u2)
	fmt.Println(u2)
}
