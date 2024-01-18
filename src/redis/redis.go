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

func Set(ctx context.Context, client *redis.Client, key string, args ...interface{}){
	err := client.HSet(ctx, key, args).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("hset was successful.")
}

func Getall(ctx context.Context, client *redis.Client, key string){
	val, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	for keys_inside, value := range val {
		fmt.Printf("%s ] %s ] %s : %s\n" ,keys_inside, common.Current_date_for_message(), key, value)
	}
	fmt.Println("hget all was successful.")
}

func Delete(ctx context.Context, client *redis.Client, key string, id string){
	cmd := client.HDel(ctx, key, id)
	if err := cmd.Err(); err != nil {
		panic(err)
	}
	if cmd.Val() <= 0 {
		fmt.Println("the id or key", key, id, "was not found")
	} else {
		fmt.Println("hdelete was successful.")
	}
	println("Fields deleted: ", cmd.Val())
	
}


func Amount(ctx context.Context, client *redis.Client, key string){
	fields, err := client.HKeys(ctx, key).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Number of keys for %s: %d\n", key, len(fields))
}


func Get(ctx context.Context, client *redis.Client, id string){
	iter := client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		result, err := client.HGet(ctx, iter.Val(), id).Result()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(iter.Val(), result)
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}