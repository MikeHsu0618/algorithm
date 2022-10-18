package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
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

	_, err = client.Collection("mts-local").Doc("LA").Set(ctx, map[string]interface{}{
		"name":    "Los Angeles",
		"state":   "CA",
		"country": "USA",
	})
	fmt.Println("done")
}
