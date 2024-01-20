package main

import (
	"main/common"
)


func main() {
	password := []byte("1qaz2wsx")
	argon2IDHash := common.HnSObject()
	hashSalt := common.HnSGenerate(password, argon2IDHash)
	common.HnSCompare(argon2IDHash, hashSalt.Hash, hashSalt.Salt, password)
}