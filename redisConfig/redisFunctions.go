package redisConfig

import (
	"Redis_Example_Using_Go/models"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

// create new redis client using go redis library
var client = redis.NewClient(&redis.Options{
	//set address
	Addr: "localhost:6379",
})

func RedisSetValues() {
	//set some values to the user struct
	json, err := json.Marshal(models.User{Name: "User", Age: 22, Address: "Galle"})
	if err != nil {
		fmt.Println(err)
	}

	//set above json object into a new redis client
	err = client.Set("1", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func RedisGetValue() {
	//get stored redis client
	val, err := client.Get("1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
