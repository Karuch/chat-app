package main

import (
	"fmt"
	"main/auth"
	//"main/server"
	"main/jwtHandler"
	"main/common"
)





func main(){
	common.ENVinit()
	//auth.Check_half_life_refresh_need_new("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY2ODI4ODUsImlhdCI6MTcwNjUxMDA4NSwic3ViIjoibGlvciJ9.XqJ2Kh33mCHvHgBiVtXYTbkrCyFSLIaSqRdPz3bVyEU")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY2OTQ2NTksImlhdCI6MTcwNjUyMTg1OSwic3ViIjoibGlvciJ9.aB2NvDE8QK0HjLkLjyEGp3QP1p2FV4wEGUrRO3wjhtI"
	
	parsedToken, err := jwtHandler.ParseRefreshToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsedToken.Subject)



	newToken, err := jwtHandler.NewRefreshToken(auth.Refresh_claim_creator(parsedToken.Subject))
	if err != nil { //if happen there's a problem with access claim creator or newaccesstoken the above checker should be valid if it got to here
		fmt.Println(err)
	}

	fmt.Println(newToken)
	//server.Start()
}