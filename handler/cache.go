package handler

import (
	"crypto/sha1"
	"gbEngine/admin"
	"gbEngine/utils"

	"github.com/go-redis/redis"
)

type CacheHandler struct {
	RedisClient *redis.Client
	Logger      *admin.Logger
}

func (cache *CacheHandler) GetUserConnectNode(id string) (string, error) {
	sha := sha1.New()
	_, err := sha.Write([]byte(id))
	if err != nil {
		return "", err
	}

	hash := sha.Sum(nil)
	b64Uuid := utils.Encode(hash)
	result := cache.RedisClient.Get("CD_" + b64Uuid)
	if result.Err() != nil {
		return "", result.Err()
	}

	return result.Val(), nil
}
