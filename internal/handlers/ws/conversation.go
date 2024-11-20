package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"go-live-chat/internal/handlers/dto"
	"slices"
	"sync"
)

type ConversationWSUseCase struct {
	Conn    *websocket.Conn
	Channel string
	Rdb     *redis.Client
	Mu      sync.Mutex
}

var allowedTypes = []string{"user"}

func (c *ConversationWSUseCase) ListenAndForward(ctx context.Context) {
	pubsub := c.Rdb.Subscribe(ctx, c.Channel)

	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
			fmt.Println("pubsub close:", err)
		}
	}(pubsub)

	fmt.Printf("Client subscribed to Redis channel: %s\n", c.Channel)

	for msg := range pubsub.Channel() {
		c.Mu.Lock()
		err := c.Conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		c.Mu.Unlock()
		if err != nil {
			fmt.Println("Error writing to WebSocket:", err)
			break
		}
	}
}

func (c *ConversationWSUseCase) PublishFromWebSocket(ctx context.Context) {
	for {

		var message dto.ConversationRequestDto
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err) {
				fmt.Printf("WebSocket unexpectedly closed: %v\n", err)
			}

		}

		if !slices.Contains(allowedTypes, message.Type) {
			_ = c.Conn.WriteMessage(websocket.TextMessage, []byte("Not Allowed Type"))
		}

		if jsonData, err := json.Marshal(message); err == nil {
			err = c.Rdb.Publish(ctx, message.Destination, jsonData).Err()
			if err != nil {
				fmt.Printf("Failed to publish to Redis channel %s: %v\n", message.Destination, err)
				break
			}

			fmt.Printf("Message published to channel %s: %s\n", message.Destination, message)
		} else {
			fmt.Printf("Failed to publish to Redis channel %s: %v\n", message.Destination, err)
		}

	}
}

func (c *ConversationWSUseCase) Close() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("pubsub close:", err)
		return
	}
	fmt.Printf("Client disconnected from channel: %s\n", c.Channel)
}
