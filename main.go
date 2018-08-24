package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ivahaev/effectiveoctohappiness/redis"
	"github.com/ivahaev/effectiveoctohappiness/server"
)

var (
	// Default HTTP port.
	port = "3000"

	redisAddr = "localhost:6379"
	redisPass = ""
	redisDB   = 0
)

func main() {
	if v := os.Getenv("EFFECTIVE_HTTP_PORT"); len(v) > 0 {
		port = v
	}
	if v := os.Getenv("EFFECTIVE_REDIS_ADDR"); len(v) > 0 {
		redisAddr = v
	}
	if v := os.Getenv("EFFECTIVE_REDIS_PASS"); len(v) > 0 {
		redisPass = v
	}
	if v := os.Getenv("EFFECTIVE_REDIS_DB"); len(v) > 0 {
		id, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Invalid EFFECTIVE_REDIS_DB value %q, parsing error: %v", v, err)
		}
		redisDB = id
	}

	redis.Connect(redisAddr, redisPass, redisDB)
	server.Start(port)
}
