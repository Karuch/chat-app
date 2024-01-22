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


type Client_login struct {
	User string `json:"user"`
	Password string `json:"password"`
}


type Client_response struct {
	Body string `json:"Body"`
	Authorization string `json:"Authorization"`
}


func response_handler(c *gin.Context, ) Client_login {
	ResponsejsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
	}
	
	var client_login Client_login

	err = json.Unmarshal(ResponsejsonData, &client_login)
	if err != nil {
		fmt.Println("Error:", err)
	}
	
	return client_login
}



func createUserHandler(c *gin.Context) { //user/create
	// Handle response
	c.JSON(http.StatusOK, gin.H{
		"status": "Request sent successfully",
		"token": "token",
	})
	fmt.Println(auth.Create_user(postgres.Client_connect(), response_handler(c).User, response_handler(c).Password))
	fmt.Println(postgres.Get_all_messages(postgres.Client_connect(), "guyz"))
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
