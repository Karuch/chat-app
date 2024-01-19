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
  fmt.Println(redis.Delete(redis.Ctx, redis.Connect_to_db(), "5491590d-e568-44d0-8c7d-c6785ebf7e2a"))
  fmt.Println(redis.Set(redis.Ctx, redis.Connect_to_db(), "tal", "hello redis"))
  fmt.Println(redis.Getall(redis.Ctx, redis.Connect_to_db(), "tal"))
  fmt.Println(redis.Get(redis.Ctx, redis.Connect_to_db(), "cdbe806a-9357-47cf-a6bf-a423611eb710"))
  
}

