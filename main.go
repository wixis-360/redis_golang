package main

import "Redis_Example_Using_Go/redisConfig"

func main() {
	//call set value method
	redisConfig.RedisSetValues()
	//call get value method
	redisConfig.RedisGetValue()
}
