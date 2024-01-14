package main

import (
	"context"
	"log"
	"time"

	"github.com/eduardylopes/go-microservices/logger-service/data"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, res *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*res = "logged via RPC"
	return nil
}
