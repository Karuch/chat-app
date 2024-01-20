package main

import (
	"fmt"
	"main/common"
	"main/jwtHandler"
	"time"

	"github.com/golang-jwt/jwt"
	//"main/postgres"
)


func main() {
	common.ENVinit()
  
    defer func() {
        if r := recover(); r != nil {
        	common.CustomErrLog.Println("Recovered from PANIC", r)
        }
    }()

	// user login validation should occur here
  
	userClaims := jwtHandler.UserClaims{
		Username: "Leeroy",
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
	
	check_access_token(access)
	check_refresh_token(refresh)

}

func check_access_token(accesstoken string){
	parsedAccessToken := jwtHandler.ParseAccessToken(accesstoken)
	if jwtHandler.Validate_access(parsedAccessToken) {
		//allow user use it's user name to do stuff
	} else {
		//ask user for refresh token then check it
	}
}

func check_refresh_token(refreshtoken string){
	parsedRefreshToken := jwtHandler.ParseRefreshToken(refreshtoken)
	if jwtHandler.Validate_refresh(parsedRefreshToken) {
		//send user accesstoken then check if it's valid
	} else {
		//ask user to login again
	}
}
