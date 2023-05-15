package main

import (
	"gbEngine/admin"
	"gbEngine/config"
	"gbEngine/controllers/controller_v1"
	"gbEngine/handler"
	"gbEngine/routes"
	"log"
)

// / struct that contains all environment variables
var ENV *config.ENV

// a remote logging mechanism
var LOGGER *admin.Logger

// a struct which has the access to the all the resources that is
// required for this engine to work properly
var HANDLERS handler.Handler

func init() {

	var err error

	// LOADING ENVIRONMENT VARIABLES
	ENV = config.LoadENV()

	LOGGER, err = admin.InitializeLogger(ENV, "engine")
	if err != nil {
		log.Fatal("[LOGGER ERROR] : ", err)
	}

	db, err := config.MongoDBConnect(ENV)
	if err != nil {
		log.Fatal("[MONGODB ERROR] : ", err)
	}

	cache, err := config.ConnectRedis(ENV)
	if err != nil {
		log.Fatal("[REDIS ERROR] : ", err)
	}

	queue, err := config.ConnectToQueue(ENV)
	if err != nil {
		log.Fatal("[RABBITMQ ERROR] : ", err)
	}

	HANDLERS = handler.Handler{
		DataBase: &handler.DataBaseHandler{
			Mongo:  *db,
			Logger: LOGGER,
		},
		Cache: &handler.CacheHandler{
			RedisClient: cache,
			Logger:      LOGGER,
		},
		Queue: &handler.QueueHandler{
			Queue:  *queue,
			Logger: LOGGER,
		},
	}
}

func main() {
	ctrl := controller_v1.Controller{
		Handler: &HANDLERS,
		Logger:  LOGGER,
	}

	routes.HandleJob(&ctrl, LOGGER)
}
