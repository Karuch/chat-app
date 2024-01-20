package main

import (
	"fmt"
	"main/common"
	"main/jwtHandler"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	//"main/postgres"
)


func main() {
	common.ENVinit()
  
	// user login validation should occur here
  
	userClaims := jwtHandler.UserClaims{
		Id:    "01H4EKGQSY5637MQP395283JR8",
		First: "Leeroy",
		Last:  "Jenkins",
		StandardClaims: jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	
	refreshClaims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
	}
	
	access := jwtHandler.New_signed_access_token(userClaims)
	refresh := jwtHandler.New_signed_refresh_token(refreshClaims)
	
	fmt.Println(access)
	fmt.Println(refresh)
	
	jwtHandler.Validate_refresh(jwtHandler.ParseRefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU5NTYwNjYsImlhdCI6MTcwNTc4MzI2Nn0.s1j5clJx6us6Af5MExGJ8ey2bkl445UNJY_6wO7gFI8"))
	fmt.Println(os.Getenv("TOKEN_SECRET"))
	
}


