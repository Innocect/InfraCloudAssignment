package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"shortlink/internal/models"

	"github.com/go-redis/redis"
)

var client *redis.Client

// InitShortlinkRedisDAO returns pointer to instance of ShortlinkRedisDAO implemenation
func InitShortlinkRedisDAO() *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func SetShortURL(ctx context.Context, shortURL string, longURL string) error {
	redisData := &models.ShortLinkRedisData{
		LongURL:  longURL,
		ShortURL: shortURL,
	}

	marshalledRedisData, err := json.Marshal(redisData)
	if err != nil {
		fmt.Print("RedisDao Error : Unable to marshal json response")
		return err
	}

	err = client.Set(longURL, marshalledRedisData, 0).Err()
	if err != nil {
		fmt.Print("RedisDao Error : Unable to set data in redis")
		return err
	}

	return nil
}

func GetShortURL(ctx context.Context, longURL string) (string, error) {

	redisData := &models.ShortLinkRedisData{}

	unmarshalledRedisData, err := client.Get(longURL).Bytes()
	if err != nil {
		fmt.Print("RedisDao Error : Unable to retrieve data from Redis")
		return "", err
	}

	err = json.Unmarshal(unmarshalledRedisData, redisData)
	if err != nil {
		fmt.Print("RedisDao Error : Unable to unMarshal redis Data")
		return "", err
	}

	if redisData != nil {
		return redisData.ShortURL, nil
	}
	return "", nil
}
