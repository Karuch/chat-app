package main

import (
	"main/common"
	"fmt"
	"main/postgres"
)


func main() {
	common.ENVinit()
	password := []byte("1qaz2wsx")
	argon2IDHash := common.HnSObject()
	hashSalt := common.HnSGenerate(password, argon2IDHash)
	fmt.Println(common.HnSCompare(argon2IDHash, hashSalt.Hash, hashSalt.Salt, password))
	postgres.Send_db_command_test(postgres.Client_connect(), "INSERT INTO USERS (username, hash, salt) VALUES ($1, $2, $3);", "tap", hashSalt.Hash, hashSalt.Salt)
}