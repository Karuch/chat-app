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

	// Define the "/user/create" endpoint
	userGroup.POST("/create", createUserHandler)

	// Define the "/user/login" endpoint
	userGroup.POST("/login", loginUserHandler)

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}	
}	


func response_handler(c *gin.Context) map[string]interface{} {

    clientResponse := make(map[string]interface{})

    err := json.NewDecoder(c.Request.Body).Decode(&clientResponse)
    if err != nil {
        return nil
    }

    return clientResponse
}

func createUserHandler(c *gin.Context) { //user/create
	map_of_user_data := response_handler(c)

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
	map_of_user_data := response_handler(c)

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
