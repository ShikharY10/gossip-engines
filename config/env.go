package config

import (
	"crypto/sha1"
	"fmt"
	"gbEngine/utils"
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	MongoDBConnectionMethod string // manual
	MongoDBPort             string // 27017
	MongoDBHost             string // 127.0.0.1
	MongoDBUsername         string // rootuser
	MOngoDBPassword         string // rootpass
	MongoDBConnectionString string // mongodb connection string will be used when MongoDBConnectionMethod is set to auto
	RedisHost               string // 127.0.0.1
	RedisPort               string // 6379
	RabbitMQHost            string // 127.0.0.1
	RabbitMQPort            string // 5672
	RabbitMQUsername        string // guest
	RabbitMQPassword        string // guest
	EngineName              string // EN_____
	EngineMode              string // debug
	LogServerHost           string // 127.0.0.1
	LogServerPort           string // 6002
}

func LoadENV() *ENV {
	godotenv.Load()
	var env ENV

	var value string
	var found bool

	value, found = os.LookupEnv("MONGODB_CONNECTION_METHOD")
	if found {
		env.MongoDBConnectionMethod = value
	} else {
		env.MongoDBConnectionMethod = "manual"
	}

	value, found = os.LookupEnv("MONGODB_PORT")
	if found {
		env.MongoDBPort = value
	} else {
		env.MongoDBPort = "27017"
	}

	value, found = os.LookupEnv("MONGODB_HOST")
	if found {
		env.MongoDBHost = value
	} else {
		env.MongoDBHost = "127.0.0.1"
	}

	value, found = os.LookupEnv("MONGODB_USERNAME")
	if found {
		env.MongoDBUsername = value
	} else {
		env.MongoDBUsername = "rootuser"
	}

	value, found = os.LookupEnv("MONGODB_PASSWORD")
	if found {
		env.MOngoDBPassword = value
	} else {
		env.MOngoDBPassword = "rootpass"
	}

	value, found = os.LookupEnv("MONGODB_CONNECTION_STRING")
	if found {
		env.MongoDBConnectionString = value
	} else {
		env.MongoDBConnectionString = ""
	}

	value, found = os.LookupEnv("REDIS_HOST")
	if found {
		env.RedisHost = value
	} else {
		env.RedisHost = "127.0.0.1"
	}

	value, found = os.LookupEnv("REDIS_PORT")
	if found {
		env.RedisPort = value
	} else {
		env.RedisPort = "6379"
	}

	value, found = os.LookupEnv("RabbitMQHost")
	if found {
		env.RabbitMQHost = value
	} else {
		env.RabbitMQHost = "127.0.0.1"
	}

	value, found = os.LookupEnv("RABBITMQ_PORT")
	if found {
		env.RabbitMQPort = value
	} else {
		env.RabbitMQPort = "5672"
	}

	value, found = os.LookupEnv("RABBITMQ_USERNAME")
	if found {
		env.RabbitMQUsername = value
	} else {
		env.RabbitMQUsername = "guest"
	}

	value, found = os.LookupEnv("RABBITMQ_PASSWORD")
	if found {
		env.RabbitMQPassword = value
	} else {
		env.RabbitMQPassword = "guest"
	}

	value, found = os.LookupEnv("ENGINE_NAME")
	if found {
		env.EngineName = value
	} else {
		hash := sha1.New()
		hash.Write(utils.GenerateAesKey(10))
		hashDigest := hash.Sum(nil)

		hashString := fmt.Sprintf("%x", hashDigest)
		env.EngineName = "EN_" + hashString
	}

	value, found = os.LookupEnv("ENGINE_MODE")
	if found {
		env.EngineMode = value
	} else {
		env.EngineMode = "debug"
	}

	value, found = os.LookupEnv("LOG_SERVER_HOST")
	if found {
		env.LogServerHost = value
	} else {
		env.LogServerHost = "127.0.0.1"
	}

	value, found = os.LookupEnv("LOG_SERVER_PORT")
	if found {
		env.LogServerPort = value
	} else {
		env.LogServerPort = "10223"
	}

	return &env
}
