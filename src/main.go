package main

import (
	"fmt"
	"main/auth"
	"main/common"
	
	"main/postgres"


	
	//"main/postgres"
)


func main() {
	common.ENVinit()
  
    defer func() {
        if r := recover(); r != nil {
        	common.CustomErrLog.Println("Recovered from PANIC", r)
        }
    }()

	fmt.Println(auth.Create_user(postgres.Client_connect(), "maor", "1qaz3edc"))
	fmt.Println(auth.Validate_user(postgres.Client_connect(), "maor", "1qaz3edc"))

}