package main

import (
  //"main/postgres"
  "main/redis"
  "main/common"
)

func main() {
  common.ENVinit()
  redis.Amount(redis.Ctx, redis.Connect_to_db(), "elad")
}

