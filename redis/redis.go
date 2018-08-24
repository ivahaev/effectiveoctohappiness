package redis

import (
	"log"

	"github.com/go-redis/redis"
)

var client *redis.Client

// Connect creates Redis DB client. It panics when can't ping DB.
func Connect(addr, pass string, db int) {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	_, err := c.Ping().Result()
	if err != nil {
		log.Fatalf("Connection to Redis %q failed, error: %v", addr, err)
	}

	client = c
}

// Set sets value for key in DB with no expiration.
func Set(key, value string) error {
	return client.Set(key, value, 0).Err()
}
