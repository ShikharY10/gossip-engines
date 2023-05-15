package handler

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"gbEngine/admin"
	"gbEngine/config"
	"gbEngine/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBaseHandler struct {
	Mongo  config.MongoDB
	Logger *admin.Logger
}

func (db *DataBaseHandler) GetUsersData(filter bson.M, findOption *options.FindOptions) (*[]models.User, error) {
	cursor, err := db.Mongo.Users.Find(context.TODO(), filter, findOption)
	if err != nil {
		return nil, err
	}

	var users []models.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	if len(users) > 0 {
		return &users, nil
	}
	return nil, errors.New("no document found")
}

func (db *DataBaseHandler) GetUserData(filter bson.M, findOptions *options.FindOneOptions) (*models.User, error) {
	cursor := db.Mongo.Users.FindOne(context.TODO(), filter, findOptions)

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	var user models.User
	err := cursor.Decode(&user)
	if err != nil {
		return nil, err
	} else {

		return &user, nil
	}
}

func (db *DataBaseHandler) GetUserDeliveryId(id string) (*primitive.ObjectID, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	opts := options.FindOne().SetProjection(bson.D{
		{Key: "deliveryId", Value: 1},
	})
	user, err := db.GetUserData(
		bson.M{"_id": _id},
		opts,
	)
	if err != nil {
		return nil, err
	}

	return &user.DeliveryId, nil
}

func (db *DataBaseHandler) GetUserAllDeliveryPacket(deliveryId string) ([]string, error) {
	return []string{}, nil
}

func (db *DataBaseHandler) SaveDeliveryPacket(deliveryId primitive.ObjectID, packetId string, packet []byte) error {
	fmt.Println("_id: ", deliveryId)

	filter := bson.M{"_id": deliveryId}

	result, err := db.Mongo.Payloads.UpdateOne(
		context.TODO(),
		filter,
		bson.M{"$set": bson.M{"packets." + packetId: hex.EncodeToString(packet)}},
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount > int64(0) {
		return nil
	}
	return errors.New("error storing packet to database")
}

func (db *DataBaseHandler) RemoveDeliveryPacket(deliveryId string, packetId string) error {
	_id, err := primitive.ObjectIDFromHex(deliveryId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id}

	result, err := db.Mongo.Payloads.UpdateOne(
		context.TODO(),
		filter,
		bson.M{"$unset": bson.M{"packets." + packetId: ""}},
	)

	if err != nil {
		return err
	}

	if result.ModifiedCount > int64(0) {
		return nil
	}

	return errors.New("error while removing packet")
}

func (db *DataBaseHandler) RemoveMakePartnerRequest(elementName string, id string, requestId string) error {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id}
	update := bson.M{"$unset": bson.M{(elementName + "." + requestId): ""}}

	result, err := db.Mongo.Users.UpdateOne(
		context.TODO(),
		filter,
		update,
	)

	if err != nil {
		return err
	}

	if result.ModifiedCount > int64(0) {
		return nil
	}

	return errors.New("error while removing partner request")
}

func (db *DataBaseHandler) RemovePartner(fromId string, partnerId string) error {
	_id, err := primitive.ObjectIDFromHex(fromId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id}

	result, err := db.Mongo.Users.UpdateOne(
		context.TODO(),
		filter,
		bson.M{"$pull": bson.M{"partners": partnerId}},
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount > int64(0) {
		return nil
	} else {
		return errors.New("error while updating partner")
	}
}

func (db *DataBaseHandler) GetAllDeliveryPackets(deliveryId primitive.ObjectID) (map[string]string, error) {
	filter := bson.M{"_id": deliveryId}
	opts := options.FindOne().SetProjection(bson.D{
		{Key: "packets", Value: 1},
	})

	type Delivery struct {
		ID      primitive.ObjectID `bson:"_id" json:"_id"`
		Packets map[string]string  `bson:"packets" json:"packets"`
	}

	cursor := db.Mongo.Payloads.FindOne(
		context.TODO(),
		filter,
		opts,
	)

	var delivery Delivery
	err := cursor.Decode(&delivery)
	if err != nil {
		return map[string]string{}, err
	}
	return delivery.Packets, nil
}
