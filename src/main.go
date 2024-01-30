package main

import (
	"main/common"
	"main/postgres"
	"main/redis"
	"main/server"
)





func main(){
	common.ENVinit()
	redis.Client_connect()
	postgres.Client_connect()
	server.Start()
}

