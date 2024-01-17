package main

import (
  //"main/postgres"
  "main/redis"
)

func main() {
  redis.Amount(redis.Ctx, redis.Connect_to_db(), "elad")
}

