package server

import (
	"main/auth"
	"main/common"
	"main/jwtHandler"
	"encoding/json"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

func RespBodyHandler(c *gin.Context) (map[string]interface{}, error) {

    clientResponse := make(map[string]interface{})

    err := json.NewDecoder(c.Request.Body).Decode(&clientResponse) //the function to create something from body is being called anyway
    if err != nil {													//but there's no body in normal getall request
        return nil, errors.New("request probably have no body at all, map is nil couldn't use json in map")
    }
	
    return clientResponse, nil
}


func tokenRecognizer(c *gin.Context) (bool, map[string]interface{}, error) { //this function return true only if access works even if ref work it won't true
	var token string = c.GetHeader("token")
	var tokenType string = c.GetHeader("tokenType") 
	if tokenType == "refresh" {

		parsedToken, err := jwtHandler.ParseRefreshToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "refresh_is_wrong",
				"body": "failure: token is invalid, try login",
			})
			return false, nil, err
		}
		
		newToken, err := jwtHandler.NewAccessToken(auth.AccessClaimCreator(parsedToken.Subject))
		if err != nil { //if happen there's a problem with access claim creator or newaccesstoken the above checker should be valid if it got to here
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "refresh_is_wrong",
				"body": "failue: an unknown error occurred, try again",
			})
			return false, nil, err
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "refresh_is_true",
			"body": newToken,
		})
		return false, nil, nil
		
	} else if tokenType == "access" {

		parsedToken, err := jwtHandler.ParseAccessToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "access_is_wrong",
				"body": "failure: token is invalid or expire, try refresh",
			})
			return false, nil, err
		}

		respBody, err := RespBodyHandler(c) //mmm wonder if I should check if nil or not
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{				
				"status": "access_is_true",					
				"body": "error: access token is valid but invalid request - request must have body, even empty",
			})
			return false, nil, err
		}

		respBody["username"] = parsedToken.Username

		return true, respBody, nil

	} else {

		common.CustomErrLog.Println("if it gots here it probably means that client changed manually tokentype")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "access_is_wrong", //umm probably change this later
			"body": "failure: invalid request",
		})
		return false, nil, errors.New("tokenType of 'access' or 'refresh' was not found")

	}
}