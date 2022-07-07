package redisAction

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func (r *Redis) Init(redisIP string) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisIP + ":6379",
		Password: "",
		DB:       0,
	})
	s := client.Ping()
	fmt.Println(s.String())
	r.Client = client
}

func (r *Redis) SetEngineName(name string) {
	res := r.Client.LPush("engines", name)
	if res.Err() != nil {
		log.Println(res.Err().Error())
	}
}

func (r *Redis) GetEngineName() []string {
	ress := r.Client.LRange("engines", 0, -1)
	engines, err := ress.Result()
	if err != nil {
		log.Println(err.Error())
	}
	return engines
}
