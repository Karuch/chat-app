package main

import (
	"main/common"
	"main/postgres"
	"fmt"
)


func main() {
	common.ENVinit()
	text, is_user_valid := postgres.Validate_user(postgres.Client_connect(), "taa", "1qaz2wwsx")
	fmt.Println(text, is_user_valid)
	fmt.Println(postgres.Create_user(postgres.Client_connect(), "taa", "1qaz2wwsx"))
}