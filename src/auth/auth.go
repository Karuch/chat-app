package auth

import (
	"database/sql"
	"fmt"
	"main/common"
	"github.com/lib/pq"
	"main/jwtHandler"
	
	"time"
)


func Create_user(db *sql.DB, username string, password string) string{
	hashSalt := HnSGenerate([]byte(password), ArgonObject)

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

func Validate_userpass(db *sql.DB, username string, password string) {
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
	text, user_is_valid := HnSCompare(ArgonObject, db_hash, db_salt, []byte(password))
	fmt.Println(text)
	if user_is_valid {
		//send user refresh with status login_is_true
	} else {
		//send user with status login_is_wrong
	}
	return 
}

func Check_access_token(accesstoken string){
	parsedAccessToken := jwtHandler.ParseAccessToken(accesstoken)
	if jwtHandler.Validate_access(parsedAccessToken) {
		//allow user use it's user name to do stuff with status access_is_true
	} else {
		//ask user for refresh token then check it with status access_is_wrong
	}
	return
}

func Check_refresh_token(refreshtoken string){
	parsedRefreshToken := jwtHandler.ParseRefreshToken(refreshtoken)
	fmt.Println(parsedRefreshToken)
	if jwtHandler.Validate_refresh(parsedRefreshToken) {
		//send user accesstoken then check if it's with status refresh_is_true
	} else {
		//ask user to login again with status refresh_is_wrong
	}
	return
}

func Check_half_life_refresh_need_new(refreshtoken string){
	parsedRefreshToken := jwtHandler.ParseRefreshToken(refreshtoken)
	expiresAtTime := time.Unix(parsedRefreshToken.ExpiresAt, 0)
	secondsDifference := time.Until(expiresAtTime).Seconds()
	if float64(common.Refresh_exp_min*60/2) < secondsDifference/2 {
		fmt.Println("reach")
		//generate new refresh send status to user half_time_refresh
	}
}

