package main

import (
  //"main/postgres"
  "context"
  "fmt"
  "github.com/redis/go-redis/v9"
  "time"
  "strconv"
  "github.com/google/uuid"
)

func current_date_for_message() string {
	currentTime := time.Now()
	return strconv.Itoa(currentTime.Year()) + "/" +
		strconv.Itoa(int(currentTime.Month())) + "/" +
		strconv.Itoa(currentTime.Day()) + " " +
		strconv.Itoa(currentTime.Hour()) + ":" +
		strconv.Itoa(currentTime.Minute()) + ":" + // Change from Hour() to Minute()
		strconv.Itoa(currentTime.Second())
}

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

  //get("foo", connect_to_db(), ctx)
  //getall(connect_to_db(), ctx)
  testset(ctx, connect_to_db(), "elad", uuid.New().String(), "testiiii")
  testdel(ctx, connect_to_db(), "elad", "0")
  testgetall(ctx, connect_to_db(), "elad")
  amount(ctx, connect_to_db(), "elad")
}

func testset(ctx context.Context, client *redis.Client, key string, args ...interface{}){
  err := client.HSet(ctx, key, args).Err()
  if err != nil {
      panic(err)
  }
  fmt.Println("hset was successful.")
}

func testgetall(ctx context.Context, client *redis.Client, key string){
  val, err := client.HGetAll(ctx, key).Result()
  if err != nil {
      panic(err)
  }
  for keys_inside, value := range val {
    fmt.Printf("%s] %s] %s: %s\n" ,keys_inside, current_date_for_message(), key, value)
  }
  fmt.Println("hget all was successful.")
}

func testdel(ctx context.Context, client *redis.Client, key string, id string){
  cmd := client.HDel(ctx, key, id)
  if err := cmd.Err(); err != nil {
      panic(err)
  }
  println("Fields deleted: ", cmd.Val())
  fmt.Println("hdelete was successful.")
}


func amount(ctx context.Context, client *redis.Client, key string){
	fields, err := client.HKeys(ctx, key).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Number of keys for %s: %d\n", key, len(fields))
}