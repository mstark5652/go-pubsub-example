package internal

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func createTopicIfNotExists(client *pubsub.Client, topicID string) *pubsub.Topic {
	ctx := context.Background()
	topic := client.Topic(topicID)
	ok, err := topic.Exists(ctx)
	if err != nil {
		_ = fmt.Errorf("failed to check if topic exists: %v", err)
	}
	if ok {
		return topic
	}
	topic, err = client.CreateTopic(ctx, topicID)
	if err != nil {
		_ = fmt.Errorf("failed to create topic for: %s\nError: %v", topicID, err)
	}
	return topic
}

// Publish message to pub/sub
func Publish(projectID, topicID, message string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient error: %v", err)
	}

	topic := createTopicIfNotExists(client, topicID)
	fmt.Printf("Publishing to topic: %s\n", topicID)

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})

	// Block unit the result is returned
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get error: %v", err)
	}
	fmt.Printf("Published a message; messageId: %v\n", id)
	return nil
}
