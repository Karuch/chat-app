package common

import (
	
	"log"
	"os"
	"strconv"
	"time"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"github.com/gin-gonic/gin"
)


var CustomErrLog = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

func ENVinit(){
	os.Setenv("TOKEN_SECRET_ACCESS", 		"access")
	os.Setenv("TOKEN_SECRET_REFRESH", 		"refresh")
	os.Setenv("ACCESS_EXP_MIN", 		"15")
	os.Setenv("REFRESH_EXP_MIN", 		"2880")
	os.Setenv("REDIS_IP", 		"172.17.0.2")
	os.Setenv("REDIS_PORT", 	"6379")
	os.Setenv("REDIS_PASSWORD", "1598")
	os.Setenv("REDIS_DB", 		"0") //0 means default DB, there's problem to store int will leave it for later
	
	os.Setenv("POSTGRES_IP",		"172.17.0.2")
	os.Setenv("POSTGRES_PORT",		"5432")
	os.Setenv("POSTGRES_USER",		"postgres")
	os.Setenv("POSTGRES_PASSWORD",	"1598")
	os.Setenv("POSTGRES_DB",		"postgres")
	fmt.Println("here", os.Getenv("POSTGRES_PORT"))
	EnvVarDeclare()

}

var (
	Refresh_exp_min int
	Access_exp_min int
)

func EnvVarDeclare(){
	Refresh_exp_min = Convert_to_int(os.Getenv("REFRESH_EXP_MIN"))
	Access_exp_min = Convert_to_int(os.Getenv("ACCESS_EXP_MIN"))
}

func Current_date_for_message() string {
	currentTime := time.Now()
	return strconv.Itoa(currentTime.Year()) + "/" +
		strconv.Itoa(int(currentTime.Month())) + "/" +
		strconv.Itoa(currentTime.Day()) + " " +
		strconv.Itoa(currentTime.Hour()) + ":" +
		strconv.Itoa(currentTime.Minute()) + ":" + // Change from Hour() to Minute()
		strconv.Itoa(currentTime.Second())
}

func Random_uuid() string{
	return uuid.New().String()
}

func Convert_to_int(str string) int {
	val, err := strconv.Atoi(str);
	if err != nil {
		CustomErrLog.Println(err)
		//panic(err)
	} else {
		return val
	}
	return val
}

/*
// CustomError is a custom type that implements the error interface.
type CserverSideErr struct {	//those are errors client should not be aware, the msg will be returned as unknown error
    Message string				//the status will be 500 server fault
}

// Error returns the error message for CustomError.
func (e *CserverSideErr) Error() string {
    return e.Message
}

// Function that returns an instance of CustomError.
func ServerSideCustom_Error(str string) error {
    return &CserverSideErr{Message: str}
}
*/


func RespBodyIsFatalField(c *gin.Context, respBody map[string]interface{}, field string) (string, error) { //check if client changed field of token
																										   //if serv get it as valid token and reach this code then secret was found!
	value, ok := respBody[field].(string)																   //usually will be used for username validation but maybe another cases in the future
	if !ok {
		CustomErrLog.Println("FATAL client try to change jwt", field, "field so it will store something else which means he have the jwt pass")
		c.JSON(http.StatusBadRequest, gin.H{ 
			"status": "access_is_true",
			"body":   ErrBadRequest.Error(),
		})
		return "", ErrBadRequest
		
	}
	if field == "username" && len(value) < 3 {
		CustomErrLog.Println("FATAL client jwt not include username field but pass validation possible means that client got access secret also somehow stored less 3 char name?")
		c.JSON(http.StatusBadRequest, gin.H{ 
			"status": "access_is_true",
			"body":   ErrBadRequest.Error(),
		})
		return "", ErrBadRequest
	}
	return value, nil

}

func RespBodyIsWrongField(c *gin.Context, respBody map[string]interface{}, field string) (string, error) { //check if client changed field of header
																										   //by default cause panic behavior but want to handle
	field, ok := respBody[field].(string)
	if !ok {
		CustomErrLog.Println("wrong field in request, maybe client tried to change manually?", field)
		c.JSON(http.StatusBadRequest, gin.H{ 
			"status": "access_is_true",
			"body":	"error: invalid field in header",
		})
		return "", ErrBadRequest
	}
	return field, nil

}