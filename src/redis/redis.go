package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"main/common"
	"os"
)

/* examples
fmt.Println(redis.Delete(redis.Ctx, redis.Connect_to_db(), "5491590d-e568-44d0-8c7d-c6785ebf7e2a"))
fmt.Println(redis.Set(redis.Ctx, redis.Connect_to_db(), "tal", "hello redis"))
fmt.Println(redis.Getall(redis.Ctx, redis.Connect_to_db(), "tal"))
fmt.Println(redis.Get(redis.Ctx, redis.Connect_to_db(), "cdbe806a-9357-47cf-a6bf-a423611eb710")) */

var DBconnect *redis.Client

func SetupRedisConnection() {
	DBconnect = Client_connect()
}

func Client_connect() *redis.Client { //wonder if that's a good idea it means that the connection will disconnect each time
	var ctx context.Context = context.Background()
	client := redis.NewClient(&redis.Options{ //calling function I think
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_IP"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),                  // no password set
		DB:       common.Convert_to_int(os.Getenv("REDIS_DB")), // use default DB
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		common.CustomErrLog.Println(err)
		//panic(err) will cause panic if postgres is down which is not a behavior I exacly want but fatal error
	}

	return client
}

func Set(ctx context.Context, client *redis.Client, user string, message string) (string, error) {
	// Specify the key and fields
	var id string = common.Random_uuid()
	fields := map[string]interface{}{
		"date":    common.Current_date_for_message(),
		"user":    user,
		"message": message,
	}
	
	// Use the HSet method to set the values
	_, err := client.HSet(ctx, id, fields).Result()
	if err != nil {
		common.CustomErrLog.Println(err)
		return "", common.ErrInternalFailure
	}

	// Print the result (1 if a new field was created, 0 if field already existed and was updated)

	return fmt.Sprintf("message has been added: %v ] %v ] %v : %v\n", id, fields["date"], fields["user"], fields["message"]), nil
}

func Getall(ctx context.Context, client *redis.Client, username string) ([]string, error) {
	all_messages := []string{}
	iter := client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		result, err := client.HGet(ctx, iter.Val(), "user").Result()

		if err != nil {
			common.CustomErrLog.Println(err)
			return []string{""}, common.ErrInternalFailure
		}
		if result == username {
			single_message, err := Get(ctx, client, username, iter.Val())
			if err != nil {
				return []string{""}, err
			}
			all_messages = append(all_messages, single_message)
		}
	}
	if err := iter.Err(); err != nil {
		common.CustomErrLog.Println(err)
		return []string{""}, common.ErrInternalFailure
	}
	return all_messages, nil
}

func Delete(ctx context.Context, client *redis.Client, username string, key string) (string, error) {
	exists_key_check, err := client.Exists(ctx, key).Result()
	if err != nil {
		common.CustomErrLog.Println(err) //happend when redis was down for exmaple
		return "", common.ErrInternalFailure
		
	}
	if exists_key_check != 1 {
		common.CustomErrLog.Println("key was not found", key)
		return "", fmt.Errorf("%w: nothing was found using this ID `%s`", common.ErrNotFound, key)
	}
	//actual function v
	values, err := client.HMGet(ctx, key, "user").Result()
	if err != nil {
		common.CustomErrLog.Println(err)
		return "", common.ErrInternalFailure
	}

	if values[0] != username { //user does not permitted but ID was found
		common.CustomErrLog.Println("id was found but not permitted", username, key)
		return "", fmt.Errorf("%w: nothing was found using this ID `%s`", common.ErrNotFound, key)
	}

	_, err = client.Del(ctx, key).Result()
	if err != nil {
		common.CustomErrLog.Println(err)
		return "", common.ErrInternalFailure
	}
	return fmt.Sprintf("message ID '%s' has been deleted successfully.", key), nil

}

func Get(ctx context.Context, client *redis.Client, username string, key string) (string, error) {
	exists_key_check, err := client.Exists(ctx, key).Result()
	if err != nil {
		common.CustomErrLog.Println(err)
		return "", common.ErrInternalFailure
	}
	if exists_key_check != 1 {
		return "", fmt.Errorf("%w: nothing was found using this ID `%s`", common.ErrNotFound, key)
	}
	//actual function v

	fields := []string{"date", "user", "message"}
	values, err := client.HMGet(ctx, key, fields...).Result()
	if err != nil {
		common.CustomErrLog.Println(err)
		return "", common.ErrInternalFailure
	}
	if values[1] != username { //user does not permitted but ID was found
		common.CustomErrLog.Println("id was found but not permitted", username, key)
		return "", fmt.Errorf("%w: nothing was found using this ID `%s`", common.ErrNotFound, key)
	}
	return fmt.Sprintf(" %v ] %v ] %v : %v", key, values[0], values[1], values[2]), nil
}
