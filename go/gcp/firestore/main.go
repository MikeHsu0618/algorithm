package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./config/credentials/super-account.json")
	fmt.Println(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	// Use a service account
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	doc := make(map[string]interface{})
	doc["stringExample"] = "Hello world!"
	doc["booleanExample"] = true
	doc["numberExample"] = 3.14159265
	doc["dateExample"] = time.Now()
	doc["arrayExample"] = []interface{}{5, true, "hello"}
	doc["nullExample"] = nil
	doc["objectExample"] = map[string]interface{}{
		"a": 5,
		"b": true,
	}
	_, err = client.Collection("mts-local").Doc("LA").Set(ctx, doc)

	type City struct {
		Name       string   `firestore:"name,omitempty"`
		State      string   `firestore:"state,omitempty"`
		Country    string   `firestore:"country,omitempty"`
		Capital    bool     `firestore:"capital,omitempty"`
		Population int64    `firestore:"population,omitempty"`
		Regions    []string `firestore:"regions,omitempty"`
	}

	city := City{
		Name:    "Los A1231211341341343ngeles",
		Country: "USㄉˇ123A",
	}

	_, err = client.Collection("mts-local").Doc("citys").Set(ctx, city)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	fmt.Println("done")
}
