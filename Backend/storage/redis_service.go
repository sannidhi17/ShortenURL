package storage

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	redisClient *redis.Client
}

var (
	ctx          = context.Background()
	redisService = &RedisService{}
)

const cacheDuration = 300 //time is in seconds

func InitializeCache() *RedisService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	res, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing Redis: %v", err))
	}

	fmt.Println("Redis started successfully")
	fmt.Println(res)
	fmt.Println(redisService)
	redisService.redisClient = redisClient
	return redisService
}

func SetURL(original string, shortURL string) { //why is userId needed?
	err := redisService.redisClient.Set(ctx, shortURL, original, cacheDuration)
	if err != nil {
		panic(fmt.Sprintf("Failed to save to Redis | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortURL, original))
	}
}

func GetURL(shortURL string) string {
	res, err := redisService.redisClient.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		fmt.Println("Provided short URL doesn't exist in Redis") //should this be an exception?
	} else if err != nil {
		panic(fmt.Sprintf("Failed to retrieve URL from Redis | Error: %v - shortUrl: %s\n", err, shortURL))
	}

	return res
}
