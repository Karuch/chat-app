package main

import (
	//"encoding/json"
	"fmt"
	"main/common"
	"main/jwtHandler"
	"time"

	"github.com/golang-jwt/jwt"
	"main/auth"
	//"main/postgres"
)



func main() {
	common.ENVinit()
	
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
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}
	
	access := jwtHandler.New_signed_access_token(userClaims)
	refresh := jwtHandler.New_signed_refresh_token(refreshClaims)
	fmt.Println(access)
	fmt.Println(refresh)
	ref := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDYwMDQwMTksImlhdCI6MTcwNTgzMTIxOX0.unBaz6XPaPRHw4CH7FnG6D8uijT2fG-Rx4dc1sbg2i8"
	parsrefresh := jwtHandler.ParseRefreshToken(ref)
	
	
	issuedAtTime := time.Unix(parsrefresh.IssuedAt, 0)
	expiresAtTime := time.Unix(parsrefresh.ExpiresAt, 0)
	notBeforeTime := time.Unix(parsrefresh.NotBefore, 0)

	
	fmt.Println(issuedAtTime)
	fmt.Println(expiresAtTime)
	fmt.Println(notBeforeTime)
	
	//currentTime := time.Now()
	//secondsDifference := expiresAtTime.Sub(currentTime).Seconds()
	auth.Check_half_life_refresh_need_new(ref)

	
}	







/*defer func() {
	if r := recover(); r != nil {
		common.CustomErrLog.Println("Recovered from PANIC", r)
	}
}()*/