package common

import (
  "time"
  "strconv"
  "os"
  "github.com/google/uuid"
  "fmt"
)

func ENVinit(){
	os.Setenv("REDIS_IP", "172.17.0.2")
	os.Setenv("REDIS_PORT", "6379")
	os.Setenv("REDIS_PASSWORD", "1598")
	os.Setenv("REDIS_DB", "0") //0 means default DB, there's problem to store int will leave it for later
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
