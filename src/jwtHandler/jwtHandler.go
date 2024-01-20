package jwtHandler

import (
	"fmt"
	"main/common"
	"os"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
 Username string `json:"Username"`
 jwt.StandardClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func NewRefreshToken(claims jwt.StandardClaims) (string, error) {
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) *UserClaims {
    parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("TOKEN_SECRET")), nil
    })

    return parsedAccessToken.Claims.(*UserClaims)
}

func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
    parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("TOKEN_SECRET")), nil
    })
    
    return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}

func New_signed_access_token(userClaims UserClaims) string {
    newSignedAccessToken, err := NewAccessToken(userClaims)
	if err != nil {
		common.CustomErrLog.Println("error creating access token")
	}
    return newSignedAccessToken
}


func New_signed_refresh_token(refreshClaims jwt.StandardClaims) string {
    newSignedRefreshToken, err := NewRefreshToken(refreshClaims)
	if err != nil {
		common.CustomErrLog.Println("error creating access token")
	}
    return newSignedRefreshToken
}

func Validate_refresh(parsedRefreshToken *jwt.StandardClaims) bool {
    // refresh token is expired
	if parsedRefreshToken.Valid() != nil {
		fmt.Println("refresh is not valid")
        return false
	} 
    fmt.Println("refresh is valid")
    return true
}

func Validate_access(parsedAccessToken *UserClaims) bool {
	// access token is expired
	if parsedAccessToken.StandardClaims.Valid() != nil {
		fmt.Println("access is not valid")
        return false
	}
    fmt.Println("access is valid")
    return true
}