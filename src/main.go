package main

import (
	"fmt"
	
	"main/common"
	"main/postgres"
	
)


func main() {
	common.ENVinit()
	text, is_user_valid := postgres.Validate_user(postgres.Client_connect(), "taa", "1qaz2wwsx")
	fmt.Println(text, is_user_valid)
	fmt.Println(postgres.Create_user(postgres.Client_connect(), "taa", "1qaz2wwsx"))
	fmt.Println(postgres.Get_all_messages(postgres.Client_connect(), "maor"))
	fmt.Println(postgres.Add_message(postgres.Client_connect(), "tal", "hello my friends!"))


	

}


