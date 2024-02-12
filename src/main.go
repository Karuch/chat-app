package main

import (
	"main/common"
	"main/server"
	"main/postgres"
	"fmt"
	"os"
)





func main(){
	common.ENVinit()
	fmt.Println(os.Getenv("TOKEN_SECRET_ACCESS"))
	postgres.SetupPGconnection()
	server.Start()
}

