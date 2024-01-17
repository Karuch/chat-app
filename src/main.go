package main

import (
  //"main/postgres"
  "context"
  "fmt"
  "github.com/redis/go-redis/v9"
)

func main() {
  fmt.Println("test")
  client := redis.NewClient(&redis.Options{
      Addr:	  "172.17.0.2:6379",
      Password: "1598", // no password set
      DB:		  0,  // use default DB
  })

  ctx := context.Background()

  err := client.Set(ctx, "foo", "bar", 0).Err()
  if err != nil {
      panic(err)
  }

  val, err := client.Get(ctx, "foo").Result()
  if err != nil {
      panic(err)
  }
  fmt.Println("foo", val)
}

