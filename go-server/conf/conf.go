package conf

import (
	"fmt"
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
	"address": "redissearch:6380"
}
var SonicConf = map[string]string{
	"host":     os.Getenv("SONIC_HOST"),
	"port":     os.Getenv("SONIC_PORT"),
	"password": os.Getenv("SONIC_PASSWORD"),
}

func init() {
	fmt.Printf("RedisConf:%v\n", RedisConf)
	fmt.Printf("SonicConf:%v", SonicConf)
}
