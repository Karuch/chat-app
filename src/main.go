package main

import (
	"fmt"
	"main/postgres"
	//"main/redis"
	"main/common"
  "os"
)

func main() {
  common.ENVinit()
  fmt.Println(common.Convert_to_int(os.Getenv("POSTGRES_PORT")))
  postgres.Get_all_messages(postgres.Connect_to_db(), "maor")
  postgres.Add_message(postgres.Connect_to_db(), "maor", "olaaa maorrr")
  postgres.Get_message(postgres.Connect_to_db(), "62bfc1d1-c9db-44bf-aa99-2d600ba1ea34")
  postgres.Remove_message(postgres.Connect_to_db(), "62bfc1d1-c9db-44bf-aa99-2d600ba1ea34")
  //redis.Getall(redis.Ctx, redis.Connect_to_db(), "elad")
}

