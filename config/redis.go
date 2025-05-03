package config

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	// Initialize Redis connection here
	// Example:
	var user = os.Getenv("REDIS_USER")
	var pass = os.Getenv("REDIS_PASS")
	var db = os.Getenv("REDIS_DB")
	var host = os.Getenv("REDIS_HOST")
	var port = os.Getenv("REDIS_PORT")

	opt, err := redis.ParseURL("redis://" + user + ":" + pass + "@" + host + ":" + port + "/" + db)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	fmt.Println("Redis connection established")

	return client
}
