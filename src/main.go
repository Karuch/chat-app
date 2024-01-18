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

  //redis.Getall(redis.Ctx, redis.Connect_to_db(), "2")
  redis.Get(redis.Ctx, redis.Connect_to_db(), "9e959c8c-8f9c-43fd-8717-d13c7193bcc4")
}

