package server

import (
	
	"main/common"
	"main/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ShortGet(c *gin.Context) {
	haveAccess, respBody, err := tokenRecognizer(c)

	if err != nil {	//force to login state
		common.CustomErrLog.Println(err) 
		return
	}
	if haveAccess {
		message := redis.Get(c, redis.Client_connect(), respBody["username"].(string), respBody["id"].(string))
		c.JSON(http.StatusOK, gin.H{		
			"status": "access_is_true",							
			"body": message,
		})
	}
}



func ShortAdd(c *gin.Context) {
	haveAccess, respBody, err := tokenRecognizer(c)
	
	if err != nil {	//force to login state
		common.CustomErrLog.Println(err) 
		return
	}
	if haveAccess {
		result := redis.Set(c, redis.Client_connect(), respBody["username"].(string), respBody["message"].(string))
		c.JSON(http.StatusOK, gin.H{		
			"status": "access_is_true",							
			"body": result,
		})
	}
}