package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	service_email "newsletter-sub/email"
	models_email "newsletter-sub/models/email"
	"newsletter-sub/utils/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

var appConfig = config.AppCfg

func main() {
	fmt.Println("***Waiting for message...")

	// Create a context and a Pub/Sub client.
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, appConfig.PubSub.ProjectId, option.WithCredentialsFile("credential.json"))
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}

	// Get a reference to the subscription.
	subscription := client.Subscription(appConfig.PubSub.SubscriptionName)

	// Create a channel to handle received messages.
	messageChannel := make(chan *pubsub.Message)

	// Start a goroutine to receive messages from the subscription.
	go func() {
		ctx := context.Background()
		err := subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			// Process the received message.
			fmt.Printf("Received message: %s\n", string(msg.Data))

			var payload models_email.EmailPayload
			json.Unmarshal(msg.Data, &payload)

			emailService := service_email.NewEmailService()
			emailService.SendEmail(&payload)

			// Acknowledge the message to mark it as processed.
			msg.Ack()
		})
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
	}()

	// Set up a signal handler to gracefully exit the subscriber.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait for a signal to exit.
	<-stop

	// Close the message channel and stop the message processing goroutine.
	close(messageChannel)

	// Allow some time for pending acknowledgments before exiting.
	time.Sleep(5 * time.Second)
	fmt.Println("Subscriber has gracefully exited.")
}
