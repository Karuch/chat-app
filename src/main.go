package main

import (
	"fmt"
	//"main/postgres"
	"main/redis"
	"main/common"
  "os"
)

func main() {
  common.ENVinit()
  fmt.Println(common.Convert_to_int(os.Getenv("POSTGRES_PORT")))
  redis.Set(redis.Ctx, redis.Connect_to_db(), "tal", "hello redis!")

  redis.Getall(redis.Ctx, redis.Connect_to_db(), "tal")
  //redis.Get(redis.Ctx, redis.Connect_to_db(), "c3224db8-22d7-4fcc-bb4f-3f1cfa53a732")
}

