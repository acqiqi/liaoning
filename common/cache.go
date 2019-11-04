package common

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func SetCache(key string, val interface{}, expiration time.Duration) error {
	fmt.Println("setCache----->>", key)
	err := client.Set(key, val, expiration).Err()
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}

func GetCache(key string) (val string, err error) {
	val, err = client.Get(key).Result()
	if err != nil {
		//panic(err)
		return "", err
	}
	return val, nil
}

func DeleteCache(key string) (err error) {
	err = client.Del(key).Err()
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}

func SetupCache() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		//fmt.Println(pong, err)
		fmt.Println("Redis Error 报错")
		fmt.Print(err)
		fmt.Print(pong)
	}
}
