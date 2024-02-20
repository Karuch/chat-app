package server

import (
	"main/common"
	"main/redis"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ShortGet(c *gin.Context) {
	haveAccess, respBody, err := tokenRecognizer(c)

	if err != nil { //force to login state
		common.CustomErrLog.Println(err)
		return
	}

	username, err := common.RespBodyIsFatalField(c, respBody, "username")
	if err != nil {
		return
	}

	id, err := common.RespBodyIsWrongField(c, respBody, "id")
	if err != nil {
		return
	}

	if haveAccess {
		message, err := redis.Get(c, redis.DBconnect, username, id)
		if err != nil {
			common.CustomErrLog.Println(err)
			common.ErrStatusChecker(err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "access_is_true",
			"body":   message,
		})
	}
}

func ShortAdd(c *gin.Context) {
	haveAccess, respBody, err := tokenRecognizer(c)
	
	if err != nil { //force to login state
		common.CustomErrLog.Println(err)
		return
	}

	username, err := common.RespBodyIsFatalField(c, respBody, "username")
	if err != nil {
		return
	}

	message, err := common.RespBodyIsWrongField(c, respBody, "message")
	if err != nil {
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
			"body":   result,
		})
	}
}

func ShortDelete(c *gin.Context) {
	haveAccess, respBody, err := tokenRecognizer(c)
	
	if err != nil { //force to login state
		common.CustomErrLog.Println(err)
		return
	}

	username, err := common.RespBodyIsFatalField(c, respBody, "username")
	if err != nil {
		return
	}

	id, err := common.RespBodyIsWrongField(c, respBody, "id")
	if err != nil {
		return
	}

	if haveAccess {
		result, err := redis.Delete(c, redis.DBconnect, username, id)
		if err != nil {
			common.CustomErrLog.Println(err)
			common.ErrStatusChecker(err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "access_is_true",
			"body":   result,
		})
	}
}

func ShortGetall(c *gin.Context) {
	haveAccess, respBody, err := tokenRecognizer(c)

	if err != nil { //force to login state
		common.CustomErrLog.Println(err)
		return
	}

	username, err := common.RespBodyIsFatalField(c, respBody, "username")
	if err != nil {
		return
	}

	if haveAccess {
		message, err := redis.Getall(c, redis.DBconnect, username)
		if err != nil {
			common.CustomErrLog.Println(err)
			common.ErrStatusChecker(err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "access_is_true",
			"body":   message,
		})
	}
}