package main

import (
	"fmt"
	"main/common"
	"main/postgres"
	"main/redis"
	"main/server"
	"os"
)





func main(){
	common.ENVinit()
	fmt.Println(os.Getenv("TOKEN_SECRET_ACCESS"))
	postgres.SetupPGconnection()
	redis.SetupRedisConnection()
	server.Start()
}

