package common

import (
  "time"
  "strconv"
  "os"
  "github.com/google/uuid"
  "fmt"
  "log"
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
		fmt.Println("can't convert to int", err)
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