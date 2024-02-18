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
		message := redis.Get(c, redis.DBconnect, respBody["username"].(string), respBody["id"].(string))
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

	username, ok := respBody["username"].(string)
	if !ok {
		common.CustomErrLog.Println("FATAL client try to change jwt username field so it will store something else which means he have the jwt pass")
		c.JSON(http.StatusInternalServerError, gin.H{ //not really internal but client shouldn't know		
			"status": "access_is_true",							
			"body": common.ErrInternalFailure.Error(),
		})
		return
	}

	message, ok := respBody["message"].(string)
	if !ok {
		common.CustomErrLog.Println("wrong field in request, maybe client tried to change manually?")
		c.JSON(http.StatusBadRequest, gin.H{ //not really internal but client shouldn't know		
			"status": "access_is_true",							
			"body": "wrong field in request, should be 'message'",
		})
		return
	}

	if haveAccess {
		result, err := redis.Set(c, redis.DBconnect, username, message)
		if err != nil {
			common.CustomErrLog.Println(err)
			common.ErrStatusChecker(err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{		
			"status": "access_is_true",							
			"body": result,
		})
	}
}

func ShortDelete(c *gin.Context) {
	haveAccess, respBody, err := tokenRecognizer(c)

	if err != nil {	//force to login state
		common.CustomErrLog.Println(err) 
		return
	}
	if haveAccess {
		result := redis.Delete(c, redis.DBconnect, respBody["username"].(string), respBody["id"].(string))
		c.JSON(http.StatusOK, gin.H{		
			"status": "access_is_true",							
			"body": result,
		})
	}
}