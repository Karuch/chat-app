package postgres

import (
	"database/sql"
	"fmt"
	"main/common"
	"github.com/lib/pq"
)


func Create_user(db *sql.DB, username string, password string) string{
	hashSalt := common.HnSGenerate([]byte(password), common.ArgonObject)

	_, err := db.Exec("INSERT INTO USERS (username, hash, salt) VALUES ($1, $2, $3);", username, hashSalt.Hash, hashSalt.Salt)
	if err != nil {
		fmt.Println(err)
		// Analyze the PostgreSQL error
		pqErr, ok := err.(*pq.Error)
		if !ok {
			// Not a PostgreSQL error
			fmt.Println(err)
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

