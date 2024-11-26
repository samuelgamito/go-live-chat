package ws

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"go-live-chat/internal/handlers/dto"
	"log"
	"sync"
)

type ConversationWSUseCase struct {
	Conn    *websocket.Conn
	Channel string
	Rdb     RedisClient
	Mu      sync.Mutex
	UseCase map[string]ConversationUseCase
}

func (c *ConversationWSUseCase) getUseCase(scenario string) ConversationUseCase {

	uc, exists := c.UseCase[scenario]
	if !exists {
		return nil
	}

	return uc
}

func (c *ConversationWSUseCase) ListenAndForward(ctx context.Context) {
	pubsub := c.Rdb.Subscribe(ctx, c.Channel)

	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
			log.Println("pubsub close:", err)
		}
	}(pubsub)

	log.Printf("Client subscribed to Redis channel: %s\n", c.Channel)

	for msg := range pubsub.Channel() {
		c.Mu.Lock()
		err := c.Conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		c.Mu.Unlock()
		if err != nil {
			log.Println("Error writing to WebSocket:", err)
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
				log.Printf("WebSocket unexpectedly closed: %v\n", err)
			}

		}

		uc := c.getUseCase(message.Type)
		if uc == nil {
			errResp := dto.ErrorBodyDTO{
				Messages: []string{"Not Allowed Type"},
			}
			jsonData, _ := json.Marshal(errResp)
			_ = c.Conn.WriteMessage(websocket.TextMessage, jsonData)
		} else {
			members, _ := uc.FindMembers(message.Destination, ctx)
			messagesPrepared := uc.PrepareMessage(members, message.Message, message.Type)
			_ = uc.StoreMessage(messagesPrepared, ctx)
			_ = uc.PublishMessage(messagesPrepared, ctx)
		}

	}
}

func (c *ConversationWSUseCase) Close() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	err := c.Conn.Close()
	if err != nil {
		log.Println("pubsub close:", err)
		return
	}
	log.Printf("Client disconnected from channel: %s\n", c.Channel)
}
