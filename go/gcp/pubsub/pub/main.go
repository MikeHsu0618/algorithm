package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./config/credentials/super-account.json")
	fmt.Println(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	err := pullMsgs("", "")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func pullMsgs(projectID, subID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}
	return nil
}
