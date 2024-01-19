package main

import (
	//"fmt"
	//"main/postgres"
	//"main/redis"
	//"main/common"
	"fmt"
	"log"
	"main/service"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {

  // user login validation should occur here
  
  userClaims := service.UserClaims{
   Id:    "01H4EKGQSY5637MQP395283JR8",
   First: "Leeroy",
   Last:  "Jenkins",
   StandardClaims: jwt.StandardClaims{
    IssuedAt:  time.Now().Unix(),
    ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
   },
  }
  
  signedAccessToken, err := service.NewAccessToken(userClaims)
  if err != nil {
    log.Fatal("error creating access token")
  }
  
  refreshClaims := jwt.StandardClaims{
    IssuedAt:  time.Now().Unix(),
    ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
  }
  
  signedRefreshToken, err := service.NewRefreshToken(refreshClaims)
  if err != nil {
    log.Fatal("error creating refresh token")
  }
  
  // do something with access, and refresh tokens.
  // i.e. return them in an HTTP response for a successful login
  fmt.Println(signedAccessToken, "\n")
  fmt.Println(signedRefreshToken, "\n\n\n\n\nXXXXXXXX")
















  var request struct {
    AccessToken  string `json:"access_token" binding:"required"`
    RefreshToken string `json:"refresh_token" binding:"required"`
  }
  
  // request validation should occur here
  var at string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjAxSDRFS0dRU1k1NjM3TVFQMzk1MjgzSlI4IiwiZmlyc3QiOiJMZWVyb3kiLCJsYXN0IjoiSmVua2lucyIsImV4cCI6MTcwNTY5ODg4MywiaWF0IjoxNzA1Njk3OTgzfQ.VtnpbE8GtvBOxEYYQdoXUjIV6WAz4UnqWllMWuPIL14"
  var rt string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU4NjYwNDksImlhdCI6MTcwNTY5MzI0OX0.BZ1YTKGhSPRHiqB2zz8AwdoVj37W8R8h8QdWBFktrbs"

  uuserClaim := service.ParseAccessToken(at)
  rrefreshClaims := service.ParseRefreshToken(rt)
  fmt.Println(uuserClaim)
  fmt.Println(rrefreshClaims)
  // use parsed uuserClaim to validate user
  // i.e. userRepository.GetById(uuserClaim.Id)
    
  // refresh token is expired
  if rrefreshClaims.Valid() != nil {
   fmt.Println("refresh is not valid")
   request.RefreshToken, err = service.NewRefreshToken(*rrefreshClaims)
   if err != nil {
    log.Fatal("error creating refresh token")
   }
  }
  
  // access token is expired
  if uuserClaim.StandardClaims.Valid() != nil && rrefreshClaims.Valid() == nil {
    fmt.Println("access is not valid")
   request.AccessToken, err = service.NewAccessToken(*uuserClaim)
   if err != nil {
    log.Fatal("error creating access token")
   }
  }

  fmt.Println("hereR",request.RefreshToken)
  fmt.Println("hereA",request.AccessToken)
  


}

