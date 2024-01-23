package main

import (
	//"encoding/json"
	"main/auth"
	"main/common"
	"main/postgres"
	
	//"main/postgres"
	"encoding/json"
	"fmt"
	"io"
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


var Client_response map[string]interface{}

func response_handler(c *gin.Context, ) map[string]interface{} {
	ResponsejsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
	}
	
	err = json.Unmarshal(ResponsejsonData, &Client_response)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return Client_response
}

func createUserHandler(c *gin.Context) { //user/create
	
	user, is_field_valid := response_handler(c)["User"].(string)
	if !is_field_valid {
		c.JSON(http.StatusOK, gin.H{
			"status": "Error: invalid request",
		})
		common.CustomErrLog.Println("the packet don't have password field")
	}
	pass, is_field_valid := response_handler(c)["Password"].(string)
	if !is_field_valid {
		c.JSON(http.StatusOK, gin.H{
			"status": "Error: invalid request",
		})
		common.CustomErrLog.Println("the packet don't have password field")
	}
	
	str := auth.Create_user(postgres.Client_connect(), user, pass)
	// Handle response
	c.JSON(http.StatusOK, gin.H{
		"status": str,
		"token": "",
	})
	
}

func loginUserHandler(c *gin.Context) {
	// Handle the "/user/login" endpoint logic
	c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
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
