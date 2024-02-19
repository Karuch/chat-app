package server

import (
	"fmt"
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

	username, err := respBodyIsFatalField(c, respBody, "username")
	if err != nil {
		return
	}

	id, err := respBodyIsWrongField(c, respBody, "id")
	if err != nil {
		return
	}

	if haveAccess {
		message := redis.Get(c, redis.DBconnect, username, id)
		c.JSON(http.StatusOK, gin.H{
			"status": "access_is_true",
			"body":   message,
		})
	}
}

func ShortAdd(c *gin.Context) {
	haveAccess, respBody, err := tokenRecognizer(c)
	fmt.Println(respBody)
	if err != nil { //force to login state
		common.CustomErrLog.Println(err)
		return
	}

	username, err := respBodyIsFatalField(c, respBody, "username")
	if err != nil {
		return
	}

	message, err := respBodyIsWrongField(c, respBody, "message")
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

	username, err := respBodyIsFatalField(c, respBody, "username")
	if err != nil {
		return
	}

	id, err := respBodyIsWrongField(c, respBody, "id")
	if err != nil {
		return
	}

	if haveAccess {
		result := redis.Delete(c, redis.DBconnect, username, id)
		c.JSON(http.StatusOK, gin.H{
			"status": "access_is_true",
			"body":   result,
		})
	}
}





func respBodyIsFatalField(c *gin.Context, respBody map[string]interface{}, field string) (string, error) { //check if client changed field of token
																										   //if serv get it as valid token and reach this code then secret was found!
	value, ok := respBody[field].(string)																   //usually will be used for username validation but maybe another cases in the future
	if !ok {
		common.CustomErrLog.Println("FATAL client try to change jwt", field, "field so it will store something else which means he have the jwt pass")
		c.JSON(http.StatusBadRequest, gin.H{ 
			"status": "access_is_true",
			"body":   common.ErrBadRequest.Error(),
		})
		return "", common.ErrBadRequest
		
	}
	if field == "username" && len(value) < 3 {
		common.CustomErrLog.Println("FATAL client jwt not include username field but pass validation possible means that client got access secret also somehow stored less 3 char name?")
		c.JSON(http.StatusBadRequest, gin.H{ 
			"status": "access_is_true",
			"body":   common.ErrBadRequest.Error(),
		})
		return "", common.ErrBadRequest
	}
	return value, nil

}

func respBodyIsWrongField(c *gin.Context, respBody map[string]interface{}, field string) (string, error) { //check if client changed field of header
																										   //by default cause panic behavior but want to handle
	field, ok := respBody[field].(string)
	if !ok {
		common.CustomErrLog.Println("wrong field in request, maybe client tried to change manually?", field)
		c.JSON(http.StatusBadRequest, gin.H{ 
			"status": "access_is_true",
			"body":	"error: invalid field in header",
		})
		return "", common.ErrBadRequest
	}
	return field, nil

}
