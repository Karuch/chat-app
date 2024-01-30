package main

import (
	"fmt"
	"main/common"
	"main/postgres"
	"main/redis"
	"main/server"
)





func main(){
	common.ENVinit()
	fmt.Println(common.Refresh_exp_min, common.Access_exp_min)
	redis.Client_connect()
	postgres.Client_connect()
	server.Start()
}

