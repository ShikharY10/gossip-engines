package handler

import (
	"gbEngine/admin"
	"gbEngine/config"
)

type QueueHandler struct {
	Queue  config.Queue
	Logger *admin.Logger
}

func (queue *QueueHandler) Produce(nodeName string, data []byte) error {
	return nil
}
