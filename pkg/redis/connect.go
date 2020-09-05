package redis

import (
	"log"
	"os"
	"strconv"

	rds "github.com/go-redis/redis"
)

// Redis wrapper
type Redis struct {
	*rds.Client
}

// Connect to redis
func Connect() *Redis {
	db, err := strconv.Atoi(os.Getenv("RDB"))
	if err == nil {
		log.Fatal(err)
	}

	return &Redis{rds.NewClient(&rds.Options{
		Addr:     os.Getenv("RHOST") + ":6379",
		Password: os.Getenv("RPASS"),
		DB:       db, // 0 - default db
	})}
}
