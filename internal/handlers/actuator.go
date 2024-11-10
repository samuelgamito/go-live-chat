package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-live-chat/internal/infraestructure/databases"
	"go.uber.org/fx"
	"net/http"
	"time"
)

type ActuatorResponse struct {
	MongoStatus  string `json:"mongo_status"`
	ServerStatus string `json:"status"`
	RedisStatus  string `json:"redis_status"`
}

type ActuatorHandler struct {
	mongoClient databases.MongoDBClient
	redisClient databases.RedisClient
}

func NewActuatorHandler(mongoClient *databases.MongoDBClient, redisClient *databases.RedisClient) *ActuatorHandler {
	return &ActuatorHandler{
		mongoClient: *mongoClient,
		redisClient: *redisClient,
	}
}

func registerActuatorHandlers(a *ActuatorHandler, h *Handler) {
	h.Runner.Get("/health", a.healthCheck)
}

func (a *ActuatorHandler) healthCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	statusResponse := ActuatorResponse{
		MongoStatus:  "up",
		ServerStatus: "up",
		RedisStatus:  "up",
	}

	err := a.mongoClient.OpenChat.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		statusResponse.MongoStatus = "down"
	}

	err = a.redisClient.NotifyClientsRedis.Ping(ctx).Err()

	if err != nil {
		statusResponse.RedisStatus = "down"
	}

	if statusResponse.RedisStatus == "down" || statusResponse.MongoStatus == "down" {
		statusResponse.ServerStatus = "down"
		w.WriteHeader(http.StatusServiceUnavailable)

	} else {
		w.WriteHeader(http.StatusOK)
	}

	jsonData, err := json.Marshal(statusResponse)
	if err != nil {
		return
	}
	_, err = w.Write(jsonData)
	if err != nil {
		return
	}
}

var ModuleActuatorHandler = fx.Invoke(registerActuatorHandlers)
