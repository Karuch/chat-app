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
			return "", errors.New("an unknown error occurred")
		}
		


		
	}
	return fmt.Sprintf("'%s' Registered successfully.", username), nil
}

func Validate_userpass(db *sql.DB, username string, password string) (error) { //there's a problem with the way this
	rows, err := db.Query("SELECT hash, salt FROM USERS WHERE username = $1;", username)	//function is built can't
	if err != nil {																//return correct status
		common.CustomErrLog.Println(err)
		return common.ServerSideCustom_Error("an unknown error occured")
	}
	var db_hash []byte;
	var db_salt []byte;
	for rows.Next() {
		err := rows.Scan(&db_hash, &db_salt)
		if err != nil {
			common.CustomErrLog.Println(err)
			return common.ServerSideCustom_Error("an unknown error occured")
		}
	}
	_, user_is_valid := HnSCompare(ArgonObject, db_hash, db_salt, []byte(password))
	
	if user_is_valid {
		//send user refresh with status login_is_true
	} else {
		//send user with status login_is_wrong
		return errors.New("username or password invalid")
	}
	return nil
}

/*func Check_access_token(accesstoken string){
	parsedAccessToken := jwtHandler.ParseAccessToken(accesstoken)
	if jwtHandler.Validate_access(parsedAccessToken) {
		//allow user use it's user name to do stuff with status access_is_true
	} else {
		//ask user for refresh token then check it with status access_is_wrong
	}
	return
}

func Check_refresh_token(refreshtoken string){
	parsedRefreshToken, err := jwtHandler.ParseRefreshToken(refreshtoken)
	fmt.Println(parsedRefreshToken)
	if jwtHandler.Validate_refresh(parsedRefreshToken) {
		//send user accesstoken then check if it's with status refresh_is_true
	} else {
		//ask user to login again with status refresh_is_wrong
	}
	return
}*/

func Check_half_life_refresh_need_new(refreshtoken string){
	parsedRefreshToken, err := jwtHandler.ParseRefreshToken(refreshtoken)
	if err != nil {
		return
	}
	expiresAtTime := time.Unix(parsedRefreshToken.ExpiresAt, 0)
	secondsDifference := time.Until(expiresAtTime).Seconds()
	if float64(common.Refresh_exp_min*60/2) < secondsDifference/2 {
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

func AccessClaimCreator(user string) jwtHandler.UserClaims {
	var UserClaims = jwtHandler.UserClaims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(common.Access_exp_min)).Unix(),
		},
	}
	return UserClaims
}