package internal

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"
)

// Subscribe to given sub id
func Subscribe(projectID, subID string) error {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient error: %v", err)
	}
	fmt.Printf("Subscribing to subscription: %s\n", subID)

	var mu sync.Mutex
	received := 0
	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, message *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()

		fmt.Printf("Received message: %q\n", string(message.Data))
		message.Ack()
		received++
		if received == 1000 {
			cancel() // disable to not stop
		}
	})
	if err != nil {
		return fmt.Errorf("receive error: %v", err)
	}
	return nil
}
