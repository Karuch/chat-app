package server

import (
	"github.com/gin-gonic/gin"
)



func Start() {
	// Create a new Gin router
	router := gin.Default()

	// Create a route group for "/user"
	userGroup := router.Group("/user")
	longMsgGroup := router.Group("/longmsg")
	shortMsgGroup := router.Group("/shortmsg")
	// Define the "/user/create" endpoint
	userGroup.POST("/create", CreateUserHandler)

	// Define the "/user/login" endpoint
	userGroup.POST("/login", LoginUserHandler)

	longMsgGroup.GET("/getall", LongGetAll)

	longMsgGroup.GET("/get", LongGet)

	longMsgGroup.DELETE("/delete", LongDelete)

	longMsgGroup.POST("/add", LongAdd)



	shortMsgGroup.GET("/get", ShortGet)

	shortMsgGroup.POST("/add", ShortAdd)

	shortMsgGroup.DELETE("/delete", ShortDelete)

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}	
}	