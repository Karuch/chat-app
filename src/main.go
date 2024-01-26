package main

import (
	//"encoding/json"
	"fmt"
	"main/auth"
	"main/common"
	"main/jwtHandler"
	"main/postgres"

	//"main/postgres"
	"encoding/json"

	//"io"
	"net/http"

	"github.com/gin-gonic/gin"
)





func main() {
	common.ENVinit()

	// Create a new Gin router
	router := gin.Default()

	// Create a route group for "/user"
	userGroup := router.Group("/user")
	longMsgGroup := router.Group("/longmsg")

	// Define the "/user/create" endpoint
	userGroup.POST("/create", createUserHandler)

	// Define the "/user/login" endpoint
	userGroup.POST("/login", loginUserHandler)

	longMsgGroup.GET("/getall", longGetAll)

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}	
}	


func respBodyHandler(c *gin.Context) map[string]interface{} {

    clientResponse := make(map[string]interface{})

    err := json.NewDecoder(c.Request.Body).Decode(&clientResponse)
    if err != nil {
        return nil
    }

    return clientResponse
}

func createUserHandler(c *gin.Context) { //user/create
	map_of_user_data := respBodyHandler(c)

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

func loginUserHandler(c *gin.Context) {
	map_of_user_data := respBodyHandler(c)

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
	
	err := auth.Validate_userpass(postgres.Client_connect(), user, pass)
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



func longGetAll(c *gin.Context) { //longMsg/getall
	
	err, haveAccess, username := tokenRecognizer(c)
	fmt.Println(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{									//done auth success
			"body": fmt.Sprintf("You need to login"),
		})
		return
	}
	if haveAccess {
		all_messages := postgres.Get_all_messages(postgres.Client_connect(), username)
		fmt.Println(all_messages)
		c.JSON(http.StatusUnauthorized, gin.H{									//done auth success
			"body": all_messages,
		})
	}
}


func tokenRecognizer(c *gin.Context) (error, bool, string) { //this function return true only if access works even if ref work it won't true
	var token string = c.GetHeader("token")
	var tokenType string = c.GetHeader("tokenType") 
	if tokenType == "refresh" {
		parsedToken, err := jwtHandler.ParseRefreshToken(token)
		if err != nil {
			return err, false, ""
		}
		if jwtHandler.Validate_refresh(parsedToken) {
			newToken, err := jwtHandler.NewAccessToken(auth.AccessClaimCreator(parsedToken.Subject))
			if err != nil {
				common.CustomErrLog.Println(err)
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": "refresh_is_wrong",
					"body": "an unknown error occurred, try again",
				})
				return err, false, ""
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "refresh_is_true",
				"body": newToken,
			})
			return nil, false, ""
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "refresh_is_wrong",
				"body": "token is invalid, try login",
			})
			return nil, false, ""
		}
		
	} else if tokenType == "access" {
		parsedToken, err := jwtHandler.ParseAccessToken(token)
		if err != nil {
			return err, false, ""
		}
		if jwtHandler.Validate_access(parsedToken) { //need to get to this later case when user
			c.JSON(http.StatusOK, gin.H{									//done auth success
				"status": "access_is_true",
				"body": "",
			})
			return nil, true, parsedToken.Username
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "access_is_wrong",
				"body": "token is invalid, try refresh",
			})
			return nil, false, ""
		}
		

	}
	common.CustomErrLog.Println("unknown behavior of validate refresh if got here")
	return nil, false, ""
}
















/*defer func() {
	if r := recover(); r != nil {
		common.CustomErrLog.Println("Recovered from PANIC", r)
	}
}()




/*r.GET("/ping", func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	fmt.Println("here", c.GetHeader("Authorization"))
})


r.POST("/user", func(c *gin.Context) {

	// Handle response
	c.JSON(http.StatusOK, gin.H{
		"status": "Request sent successfully",
		"token": "token",
	})

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
	}
	

	var client_response Client_response

	err = json.Unmarshal(jsonData, &client_response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	client_response.Authorization = c.GetHeader("Authorization")

	logger.Println("Body:", client_response.Body)
	logger.Println("Auth:", client_response.Authorization)
})*/
