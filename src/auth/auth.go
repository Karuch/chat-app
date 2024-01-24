package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"main/common"
	"main/jwtHandler"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/lib/pq"
)

var UserClaims = jwtHandler.UserClaims{
	Username: "Leeroy",
	StandardClaims: jwt.StandardClaims{
	IssuedAt:  time.Now().Unix(),
	ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	},
}




func Create_user(db *sql.DB, username string, password string) (string, error) {
	hashSalt := HnSGenerate([]byte(password), ArgonObject)

	_, err := db.Exec("INSERT INTO USERS (username, hash, salt) VALUES ($1, $2, $3);", username, hashSalt.Hash, hashSalt.Salt)
	if err != nil {
		// Analyze the PostgreSQL error
		pqErr, ok := err.(*pq.Error)
		if !ok {
			// Not a PostgreSQL error
			common.CustomErrLog.Println(pqErr)
			return "", errors.New("an unkonwn error occurred")
		} else if pqErr.Code == "23505" {
			return "", errors.New("username is already taken")
		} else if pqErr.Code == "23514" {
			return "", errors.New("username is too short (need to be at least 3 characters)")
		} else {
			common.CustomErrLog.Println(err)
			return "", errors.New("an unknown error occured")
		}
		


		
	}
	return fmt.Sprintf("'%s' Registered successfully.", username), nil
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

func Refresh_claim_creator(user string) jwt.StandardClaims {
	var RefreshClaims = jwt.StandardClaims{
		Subject:	user,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(common.Refresh_exp_min)).Unix(),
	}
	return RefreshClaims
}
