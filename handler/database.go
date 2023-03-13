package handler

import (
	"context"
	"gbEngine/admin"
	"gbEngine/config"
	"gbEngine/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataBaseHandler struct {
	Mongo  config.MongoDB
	Logger *admin.Logger
}

func (db *DataBaseHandler) GetUserDeliveryId(id string) (*primitive.ObjectID, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := db.Mongo.Users.FindOne(
		context.TODO(),
		bson.M{"_id": _id},
	)
	var deliveryId models.DeliveryID
	err = result.Decode(&deliveryId)
	if err != nil {
		return nil, err
	}
	return &deliveryId.DeliveryId, nil
}

func (db *DataBaseHandler) GetUserAllDeliveryPacket(deliveryId string) ([]string, error) {
	return []string{}, nil
}

func (db *DataBaseHandler) SaveDeliveryPacket(deliveryId primitive.ObjectID, packetId string, payload []byte) error {
	return nil
}
