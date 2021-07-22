package conf

import (
	"log"
	"os"
)

var RedisConf = map[string]string{
	"name":    "redis",
	"type":    "tcp",
	"db":      os.Getenv("REDIS_DB"),
	"address": os.Getenv("REDIS_ADDRESS"),
	"auth":    os.Getenv("REDIS_AUTH"),
}
var RedisSearchConf = map[string]string{
	"address": os.Getenv("REDISSEARCH_ADDRESS"),
}

func init() {
	log.Println("RedisConf:", RedisConf)
	log.Println("RedisConf:", RedisSearchConf)
}
