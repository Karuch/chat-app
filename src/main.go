package main

import (
	"encoding/json"
	"fmt"
	"main/auth"
	"main/common"
	"time"
	//"main/postgres"
)

type Bird struct {
	Species     string
	Description string
	CreatedAt   time.Time
}

func main() {
	common.ENVinit()
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "createdAt": "2021-10-18T11:08:47.577Z"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird.CreatedAt)
	// {pigeon likes to perch on rocks 2021-10-18 11:08:47.577 +0000 UTC}
	fmt.Println(auth.Check_refresh_token(jwtHandler.n))
}



/*defer func() {
	if r := recover(); r != nil {
		common.CustomErrLog.Println("Recovered from PANIC", r)
	}
}()*/