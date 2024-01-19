package redis

import (
	"context"
	"fmt"
	"main/common"
	"github.com/redis/go-redis/v9"
	"os"
)

/* examples
Set(ctx, connect_to_db(), "elad", uuid.New().String(), "testiiii")
Delete(ctx, connect_to_db(), "elad", "0")
Getall(ctx, connect_to_db(), "elad")
Amount(ctx, connect_to_db(), "elad") */



var Ctx context.Context = context.Background()

func Connect_to_db() *redis.Client { //wonder if that's a good idea it means that the connection will disconnect each time
  client := redis.NewClient(&redis.Options{ //calling function I think
    Addr:	  fmt.Sprintf("%s:%s", os.Getenv("REDIS_IP"), os.Getenv("REDIS_PORT")),
    Password: os.Getenv("REDIS_PASSWORD"), // no password set
    DB:		  common.Convert_to_int(os.Getenv("REDIS_DB")),  // use default DB
  })
  return client
}

func Set(ctx context.Context, client *redis.Client, user string, message string){
	fields := map[string]interface{}{
		"date":    common.Current_date_for_message(),
		"user":  user,
		"message": message,
	}

	// Use the HSet command to set the values
	result, err := client.HSet(ctx, common.Random_uuid(), fields).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result (1 if a new field was created, 0 if field already existed and was updated)
	fmt.Println("HSET Result:", result)
}

func Getall(ctx context.Context, client *redis.Client, user string){
	iter := client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		result, err := client.HGet(ctx, iter.Val(), "user").Result()
		if err != nil {
			fmt.Println("Error:", err)
		}
		if result == user {
			Get(ctx, client, iter.Val())
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}


func Delete(ctx context.Context, client *redis.Client, key string){
	cmd := client.HDel(ctx, key)
	if err := cmd.Err(); err != nil {
		panic(err)
	}
	if cmd.Val() <= 0 {
		fmt.Println("the id", key, "was not found")
	} else {
		fmt.Println("hdelete was successful.")
	}
	println("Fields deleted: ", cmd.Val())
	
}


func Amount(ctx context.Context, client *redis.Client, user string) int {
	var count int = 0
	iter := client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		result, err := client.HGet(ctx, iter.Val(), "user").Result()
		if err != nil {
			fmt.Println("Error:", err)
		}
		if result == user {
			count++
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return count
}


func Get(ctx context.Context, client *redis.Client, key string) {
	result, err := client.HGet(ctx, key, "*").Result()
	if err != nil {
		fmt.Println("Error:", err)
		
	}
	fmt.Println(result)
}


