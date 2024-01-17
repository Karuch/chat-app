package main

import (
  //"main/postgres"
  "context"
  "fmt"
  "github.com/redis/go-redis/v9"
)

var ctx context.Context = context.Background()

func connect_to_db() *redis.Client { //wonder if that's a good idea it means that the connection will disconnect each time
  client := redis.NewClient(&redis.Options{ //calling function I think
    Addr:	  "172.17.0.2:6379",
    Password: "1598", // no password set
    DB:		  0,  // use default DB
  })
  return client
}

func main() {

  
  set("foo", "bar", connect_to_db(), ctx)
  del("tal", connect_to_db(), ctx)
  get("foo", connect_to_db(), ctx)
  getall(connect_to_db(), ctx)
  
}


func get(key string, client *redis.Client, ctx context.Context) string {
  val, err := client.Get(ctx, key).Result()
  if err != nil {
      panic(err)
  }
  fmt.Println(key, val)
  fmt.Println("get was successful.")
  return val
}

func set(key string, value string, client *redis.Client, ctx context.Context){
  err := client.Set(ctx, key, value, 0).Err()
  if err != nil {
      panic(err)
  }
  fmt.Println("set was successful.")
}

func getall(client *redis.Client, ctx context.Context) {
  // Start with a cursor of 0
  cursor := uint64(0)
  for {
      // Use SCAN to iterate over keys with a specified pattern
      keys, nextCursor, err := client.Scan(ctx, cursor, "*", 0).Result()
      if err != nil {
          panic(err)
      }
      // Process the keys
      for _, key := range keys {
          fmt.Println("key:", key, "value:", get(key, client, ctx))
      }
      // Update the cursor for the next iteration
      cursor = nextCursor
      // Check if we reached the end of the key space
      if cursor == 0 {
          break
      }
  }
}

func del(key string, client *redis.Client, ctx context.Context){
  iter := client.Scan(ctx, 0, key, 0).Iterator()

  for iter.Next(ctx) {
    key := iter.Val()
  
      d, err := client.TTL(ctx, key).Result()
      if err != nil {
          panic(err)
      }
  
      if d == -1 { // -1 means no TTL
          if err := client.Del(ctx, key).Err(); err != nil {
              panic(err)
          }
      }
  }
  
  if err := iter.Err(); err != nil {
    panic(err)
  }
}