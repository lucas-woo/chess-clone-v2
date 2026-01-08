package redisclient


import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var (
	SessionPrefix string = "session:"
)

func ConnectRedis() (*redis.Client, error) {
	redisAddr := os.Getenv("REDIS_ADDR");
	redisPassword := os.Getenv("REDIS_PASS");
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}	
	redisProtocol, err := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))
	if err != nil {
		return nil, err
	}	
	rdb := redis.NewClient(&redis.Options{
        Addr: redisAddr,
        Password: redisPassword, 
        DB: redisDB,  
				Protocol: redisProtocol,              
	})
	if rdb == nil {
		return nil, errors.New("error connecting to redis client")
	}
	fmt.Println("connected redis client")
	return rdb, nil
}
