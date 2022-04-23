package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type Dog struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type user2 struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
	Cat     Cat
	Dogs    []Dog
	Hobbies map[string]string
}

func main() {
	originUser := user2{
		ID:      "user1",
		Name:    "Mike",
		Gender:  "Male",
		Age:     0,
		Married: false,
		Cat: Cat{
			Name:   "any",
			Gender: "Female",
		},
		Dogs: []Dog{
			{
				Name:   "jay",
				Gender: "Male",
			},
		},
		Hobbies: map[string]string{
			"sports": "running, basketball",
		},
	}

	b, _ := json.Marshal(originUser)
	fmt.Println(string(b))

	// 特別好用的讓 output 的內容更好讀
	c, _ := json.MarshalIndent(originUser, "", "    ")
	fmt.Println(string(c))

	// {
	//    "id": "user1",
	//    "name": "Mike",
	//    "gender": "Male",
	//    "age": 0,
	//    "married": false,
	//    "Cat": {
	//        "name": "any",
	//        "gender": "Female"
	//    },
	//    "Dogs": [
	//        {
	//            "name": "jay",
	//            "gender": "Male"
	//        }
	//    ],
	//    "Hobbies": {
	//        "sports": "running, basketball"
	//    }
	// }
}
