package main

import (
	"main/common"
	"main/server"
	"fmt"
	"os"
)





func main(){
	common.ENVinit()
	fmt.Println(os.Getenv("TOKEN_SECRET_ACCESS"))
	server.Start()
}

