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

func Validate_user(db *sql.DB, username string, password string) (string, bool) {
	rows, err := db.Query("SELECT hash, salt FROM USERS WHERE username = $1;", username)
	if err != nil {
		fmt.Println(err)
	}
	var db_hash []byte;
	var db_salt []byte;
	for rows.Next() {
		err := rows.Scan(&db_hash, &db_salt)
			if err != nil {
	  		fmt.Println(err)

		}
	}
	text, user_is_valid := common.HnSCompare(common.ArgonObject, db_hash, db_salt, []byte(password))
	return text, user_is_valid
}

func Send_db_query_to(db *sql.DB, command string, args ...interface{}) []string {
	values := []string{}
	rows, err := db.Query(command, args...)
	  if err != nil {
	  fmt.Println(err)
	  values = append(values, "error")
		  return values
	  }
	  defer rows.Close()
  
	  for rows.Next() {
		  err := rows.Scan(&id, &message, &sender, &date)
		  if err != nil {
		fmt.Println(err)
		values = append(values, "error")
		return values
		  }
	  values = append(values, fmt.Sprintf("%v ] %v ] %v : %v", id, date, sender, message))
	  }
  
	  if err := rows.Err(); err != nil {
	  fmt.Println(err)
	  values = append(values, "error")
		  return values
	  }
	if len(values) <= 0 {
	  values = append(values, "nothing was found X_X")
	}
	return values
  }


