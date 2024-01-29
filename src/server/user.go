package server

import (
	"fmt"
	"main/auth"
	"main/common"
	"main/jwtHandler"
	"main/postgres"
	"net/http"
	"github.com/gin-gonic/gin"
)


func CreateUserHandler(c *gin.Context) { //user/create
	map_of_user_data, err := RespBodyHandler(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{				
			"status": "force_to_default",					
			"body": "error: access token is valid but invalid request - request must have body, even empty",
		})
		common.CustomErrLog.Println(err)
		return
	}

	user, is_field_valid := (map_of_user_data)["User"].(string)
	if !is_field_valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"body": "Error: invalid request",
		})
		common.CustomErrLog.Println("the packet don't have user field")
		return
	}
	pass, is_field_valid := (map_of_user_data)["Password"].(string)
	if !is_field_valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"body": "Error: invalid request",
		})
		common.CustomErrLog.Println("the packet don't have password field")
		return
	}
	
	str, err := auth.Create_user(postgres.Client_connect(), user, pass)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"body": err.Error(),
		})
		return
	}
	// Handle response

	token, err := jwtHandler.NewRefreshToken(auth.Refresh_claim_creator(user))
	if err != nil {
		common.CustomErrLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"body": "Error: Unknown",
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"status": "login_is_true",
		"body": str,
		"token": token,
	})
	
}

func LoginUserHandler(c *gin.Context) {
	map_of_user_data, err := RespBodyHandler(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{				
			"status": "force_to_default",					
			"body": "error: access token is valid but invalid request - request must have body, even empty",
		})
		common.CustomErrLog.Println(err)
		return
	}

	user, is_field_valid := (map_of_user_data)["User"].(string)
	if !is_field_valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"body": "Error: invalid request",
		})
		common.CustomErrLog.Println("the packet don't have user field")
		return 
	}
	pass, is_field_valid := (map_of_user_data)["Password"].(string)
	if !is_field_valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"body": "Error: invalid request",
		})
		common.CustomErrLog.Println("the packet don't have password field")
		return 
	}
	
	err = auth.Validate_userpass(postgres.Client_connect(), user, pass)
	if _, ok := err.(*common.CserverSideErr); ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"body": err.Error(),
		})
		return 
	} else if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": "login_is_wrong",
			"body": err.Error(),
		})
		return 
	}
	
	token, err := jwtHandler.NewRefreshToken(auth.Refresh_claim_creator(user))
	if err != nil {
		common.CustomErrLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"body": "Error: Unknown",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "login_is_true",
		"body": fmt.Sprintf("You are successfully logged in as '%s'", user),
		"token": token,
	})
}