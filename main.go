package main

import (
	"gbEngine/admin"
	"gbEngine/config"
	"gbEngine/controllers/controller_v1"
	"gbEngine/handler"
	"gbEngine/routes"
	"log"
)

func main() {
	// LOADING ENVIRONMENT VARIABLES
	ENV := config.LoadENV()

	logger, err := admin.InitializeLogger(ENV, "engine")
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.ConnectToDBs(ENV)
	if err != nil {
		logger.LogError(err)
		log.Fatal(err)
	}

	queue, err := config.ConnectToQueue(ENV)
	if err != nil {
		logger.LogError(err)
		log.Fatal(err)
	}

	handler_v1 := handler.Handler{
		DataBase: &handler.DataBaseHandler{
			Mongo:  *db.MongoDB,
			Logger: logger,
		},
		Cache: &handler.CacheHandler{
			RedisClient: db.RedisDB,
			Logger:      logger,
		},
		Queue: &handler.QueueHandler{
			Queue:  *queue,
			Logger: logger,
		},
	}

	ctrl := controller_v1.Controller{
		Handler: &handler_v1,
		Logger:  logger,
	}

	routes.HandleJob(&ctrl, logger)
}
