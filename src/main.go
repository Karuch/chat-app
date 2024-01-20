package main

import (
	"main/common"
	"main/postgres"
	"fmt"
)


func main() {
	common.ENVinit()
	fmt.Println(postgres.Create_user(postgres.Client_connect(), "taa", "1qaz2wwsx"))
}