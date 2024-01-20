package auth

import (
	"database/sql"
	"fmt"
	"main/common"
	"github.com/lib/pq"
	"main/jwtHandler"
)


func Create_user(db *sql.DB, username string, password string) string{
	hashSalt := common.HnSGenerate([]byte(password), common.ArgonObject)

	_, err := db.Exec("INSERT INTO USERS (username, hash, salt) VALUES ($1, $2, $3);", username, hashSalt.Hash, hashSalt.Salt)
	if err != nil {
		common.CustomErrLog.Println(err)
		// Analyze the PostgreSQL error
		pqErr, ok := err.(*pq.Error)
		if !ok {
			// Not a PostgreSQL error
			common.CustomErrLog.Println(err)
			return ""
		}
		// Check for unique violation (23505) and handle custom errors
		if pqErr.Code == "23505" {
			return "Username is already taken!"
		}
		// Check for check constraint violation and handle custom errors
		if pqErr.Code == "23514" {
			return "Username is too short (need to be at least 3 characters)"
		}
	}
	return fmt.Sprintf("'%s' Registered successfully.", username)
}

func Validate_user(db *sql.DB, username string, password string) (string, bool) {
	rows, err := db.Query("SELECT hash, salt FROM USERS WHERE username = $1;", username)
	if err != nil {
		common.CustomErrLog.Println(err)
	}
	var db_hash []byte;
	var db_salt []byte;
	for rows.Next() {
		err := rows.Scan(&db_hash, &db_salt)
			if err != nil {
			common.CustomErrLog.Println(err)
		}
	}
	text, user_is_valid := common.HnSCompare(common.ArgonObject, db_hash, db_salt, []byte(password))
	return text, user_is_valid
}

func Check_access_token(accesstoken string){
	parsedAccessToken := jwtHandler.ParseAccessToken(accesstoken)
	if jwtHandler.Validate_access(parsedAccessToken) {
		//allow user use it's user name to do stuff
	} else {
		//ask user for refresh token then check it
	}
}

func Check_refresh_token(refreshtoken string){
	parsedRefreshToken := jwtHandler.ParseRefreshToken(refreshtoken)
	if jwtHandler.Validate_refresh(parsedRefreshToken) {
		//send user accesstoken then check if it's valid
	} else {
		//ask user to login again
	}
}
