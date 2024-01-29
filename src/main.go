package main

import (
	"main/common"
	"main/server"
)





func main(){
	common.ENVinit()
	server.Start()
}