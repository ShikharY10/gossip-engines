package config

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBase struct {
	MongoDB *MongoDB
	RedisDB *redis.Client
}

type MongoDB struct {
	Users    *mongo.Collection
	Posts    *mongo.Collection
	Delivery *mongo.Collection
}

func ConnectToDBs(env *ENV) (*DataBase, error) {
	var db DataBase
	mongoClient, err := db.mongoDB()
	if err != nil {
		return nil, err
	}
	redisClient, err := db.redisDB()
	if err != nil {
		return nil, err
	}
	return &DataBase{
		MongoDB: mongoClient,
		RedisDB: redisClient,
	}, nil
}

func (db *DataBase) mongoDB() (*MongoDB, error) {
	mongoIP := ""
	var credential options.Credential
	credential.Username = ""
	credential.Password = ""

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoIP + ":27017").SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	var mongo MongoDB
	storage := client.Database("storage")
	mongo.Users = storage.Collection("users")
	mongo.Posts = storage.Collection("posts")
	mongo.Delivery = storage.Collection("delivery")

	return &mongo, nil
}

func (db *DataBase) redisDB() (*redis.Client, error) {
	redisIP := "127.0.0.1"
	options := redis.Options{
		Addr:     redisIP + ":6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(&options)
	ping := client.Ping()
	if ping.Err() != nil {
		return nil, ping.Err()
	}
	return client, nil
}
