package redis

import (
	"context"
	"fmt"
	"main/common"
	"os"
	"github.com/redis/go-redis/v9"
)

/* examples
  fmt.Println(redis.Delete(redis.Ctx, redis.Connect_to_db(), "5491590d-e568-44d0-8c7d-c6785ebf7e2a"))
  fmt.Println(redis.Set(redis.Ctx, redis.Connect_to_db(), "tal", "hello redis"))
  fmt.Println(redis.Getall(redis.Ctx, redis.Connect_to_db(), "tal"))
  fmt.Println(redis.Get(redis.Ctx, redis.Connect_to_db(), "cdbe806a-9357-47cf-a6bf-a423611eb710")) */



var Ctx context.Context = context.Background()

func Client_connect() *redis.Client { //wonder if that's a good idea it means that the connection will disconnect each time
  client := redis.NewClient(&redis.Options{ //calling function I think
    Addr:	  fmt.Sprintf("%s:%s", os.Getenv("REDIS_IP"), os.Getenv("REDIS_PORT")),
    Password: os.Getenv("REDIS_PASSWORD"), // no password set
    DB:		  common.Convert_to_int(os.Getenv("REDIS_DB")),  // use default DB
  })
  
  pong, err := client.Ping(Ctx).Result()
  if err != nil {
	  common.CustomErrLog.Println(err)
	  //panic(err) will cause panic if postgres is down which is not a behavior I exacly want but fatal error
  }
  fmt.Println(pong)
  return client
}

func Set(ctx context.Context, client *redis.Client, user string, message string) string {
	// Specify the key and fields
	var id string = common.Random_uuid()
	fields := map[string]interface{}{
		"date":    common.Current_date_for_message(),
		"user":  user,
		"message": message,
	}

	// Use the HSet method to set the values
	_, err := client.HSet(ctx, id, fields).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return "error"
	}

	// Print the result (1 if a new field was created, 0 if field already existed and was updated)
	
	return fmt.Sprintf("message has been added: %v ] %v ] %v : %v\n", id, fields["date"], fields["user"], fields["message"])
}


func Getall(ctx context.Context, client *redis.Client, username string) []string {
	all_messages := []string{}
	iter := client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		result, err := client.HGet(ctx, iter.Val(), "user").Result()
		
		if err != nil {
			fmt.Println("Error:", err)
		}
		if result == username {
			all_messages = append(all_messages, Get(ctx, client, username, iter.Val()))
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return all_messages
}


func Delete(ctx context.Context, client *redis.Client, key string) string {
	exists_key_check, err_key_check := client.Exists(ctx, key).Result()
	if err_key_check != nil {
		fmt.Println("Error checking key existence:", err_key_check)
		return "error"
	}
	if exists_key_check != 1 {
		return fmt.Sprintf("nothing was found X_X")
	}
	//actual function v

	_, err := client.Del(ctx, key).Result()
	if err != nil {
		return fmt.Sprintf("%s", err)
	}
	return fmt.Sprintf("message ID '%s' has been deleted successfully.", key)

}

func Get(ctx context.Context, client *redis.Client, username string, key string) string {
	exists_key_check, err_key_check := client.Exists(ctx, key).Result()
	if err_key_check != nil {
		fmt.Println("Error checking key existence:", err_key_check)
		return "error"
	}
	if exists_key_check != 1 {
		return fmt.Sprintf("nothing was found X_X") //key does not exist
	}
	//actual function v
	
	fields := []string{"date", "user", "message"}
	values, err := client.HMGet(ctx, key, fields...).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	if values[1] != username { //user does not permitted but ID was found 
		fmt.Println("not permitted") 
		return ""
	}
	return fmt.Sprintf(" %v ] %v ] %v : %v", key, values[0], values[1], values[2])
}


