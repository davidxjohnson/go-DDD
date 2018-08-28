package main

import (
	"os"
	"fmt"

	"bitbucket/dbconnectors"
	
	"github.com/go-redis/redis"
)

func connectToRedis() *redis.Client {
	redisClient := dbconnectors.ConnectToRedis(os.Getenv("REDISURL"))

	return redisClient
}

func retrieveRedisHash(redisClient *redis.Client, hash, field string) *redis.SliceCmd {
	fields := redisClient.HMGet(hash, field)

	return fields
}

func main() {
	redisClient := connectToRedis()
	//redisFields := map[string]interface{}{
	//	"hello!": "world",
	//}
	
	//redisClient.HMSet("123456", redisFields)
	fmt.Println(retrieveRedisHash(redisClient, "123456", "hello!"))
}
