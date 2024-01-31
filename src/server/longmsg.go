package server

import (
	"main/common"
	"main/postgres"
	"net/http"
	"github.com/gin-gonic/gin"
)



func LongGetAll(c *gin.Context) { //longMsg/getall

	haveAccess, respBody, err := tokenRecognizer(c)

	if err != nil {	//force to login state
		common.CustomErrLog.Println(err)  
		return
	}
	if haveAccess {
		all_messages, err := postgres.Get_all_messages(postgres.Client_connect(), respBody["username"].(string))
		if err != nil {
			common.CustomErrLog.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{	
				"status": "access_is_true",										
				"body": all_messages,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{		
			"status": "access_is_true",							
			"body": all_messages,
		})
	}
}


func LongGet(c *gin.Context) { //longMsg/get
	
	haveAccess, respBody, err := tokenRecognizer(c)
	
	if err != nil {	//force to login state
		common.CustomErrLog.Println(err) 
		return
	}
	if haveAccess {
		message, _, err := postgres.Get_message(postgres.Client_connect(), respBody["username"].(string), respBody["id"].(string))
		if err != nil {
			common.CustomErrLog.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{	
				"status": "access_is_true",										
				"body": message,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{		
			"status": "access_is_true",							
			"body": message,
		})
	}
}



func LongDelete(c *gin.Context) { //longMsg/delete
	
	haveAccess, respBody, err := tokenRecognizer(c)
	
	if err != nil {	//force to login state
		common.CustomErrLog.Println(err) 
		return
	}
	if haveAccess {
		result, err := postgres.Remove_message(postgres.Client_connect(), respBody["username"].(string), respBody["id"].(string))
		if err != nil {
			common.CustomErrLog.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{	
				"status": "access_is_true",										
				"body": result,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{	
			"status": "access_is_true",										
			"body": result,
		})
	}
}





func LongAdd(c *gin.Context) { //longMsg/add
	
	haveAccess, respBody, err := tokenRecognizer(c)
	
	if err != nil {	//force to login state
		common.CustomErrLog.Println(err) 
		return
	}
	if haveAccess {
		result, err := postgres.Add_message(postgres.Client_connect(), respBody["username"].(string), respBody["message"].(string))
		if err != nil {
			common.CustomErrLog.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{	
				"status": "access_is_true",										
				"body": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{	
			"status": "access_is_true",										
			"body": result,
		})
	}
}
