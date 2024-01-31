package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"main/common"
	_ "github.com/lib/pq"
)

/*  pg functions to implement in main for later use
  db := postgres.Client_connect()
  fmt.Println(postgres.Add_message(db, "maor", "hello this is tal"))
  fmt.Println(postgres.Remove_message(db, "125c1e50-6785-45f7-b40c-dda97e0130f8"))
  fmt.Println(postgres.Get_all_messages(db, "maor"))
  fmt.Println(postgres.Get_message(db, "125c1e50-6785-45f7-b40c-dda97e0130f8"))
  defer db.Close() 
*/

func Client_connect() *sql.DB {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  os.Getenv("POSTGRES_IP"), common.Convert_to_int(os.Getenv("POSTGRES_PORT")), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
  
  db, err := sql.Open("postgres", psqlInfo) //just intiallize db for sql conntections or something, dosen't check if DB is up
  if err != nil {                           //so will not return err if DB is down. if DB can't get SQL connection it's fatal error
    common.CustomErrLog.Println(err)                     
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    common.CustomErrLog.Println(err)
    //panic(err) will cause panic if postgres is down which is not a behavior I exacly want but fatal error
  } else {

    send_db_command_to(db, "CREATE TABLE IF NOT EXISTS MESSAGE (id VARCHAR(50), message VARCHAR(255), sender VARCHAR(50), date VARCHAR(50))")
    send_db_command_to(db, "CREATE TABLE IF NOT EXISTS USERS (id SERIAL PRIMARY KEY, username VARCHAR(50) UNIQUE CHECK (length(username) >= 3), hash BYTEA, salt BYTEA)")
    //create table if does not exist ^

    fmt.Println("Successfully connected!")
  }
  return db
}

func send_db_command_to(db *sql.DB, command string, args ...interface{}) error {
	_, err := db.Exec(command, args...)
	if err != nil {
		common.CustomErrLog.Println(err)
    return err
	}
  return nil
}

var (
  id string
  user string
  message string
  date string
)

func send_db_query_to(db *sql.DB, command string, args ...interface{}) ([]string, error) {
  values := []string{}
  rows, err := db.Query(command, args...)
	if err != nil {
    common.CustomErrLog.Println(err)
    values = append(values, "Unknown error")
		return values, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &message, &user, &date)
		if err != nil {
      common.CustomErrLog.Println(err)
      values = append(values, "Unknown error")
      return values, err
		}
    values = append(values, fmt.Sprintf("%v ] %v ] %v : %v", id, date, user, message))
	}

	if err := rows.Err(); err != nil {
    common.CustomErrLog.Println(err)
    values = append(values, "Unknown error")
		return values, err
	}
  //if len(values) <= 0 { -> return empty []string may cause nil behavior need to check this
  //  values = append(values, "nothing was found X_X")
  //}
  return values, nil
}


func Add_message(db *sql.DB, user string, message string) (string, error) {
  id = common.Random_uuid()
  date = common.Current_date_for_message()
  command := "INSERT INTO MESSAGE (id, message, sender, date) VALUES ('"+id+"', $1, $2, '"+date+"');" //look at the spceial handling ''
  err := send_db_command_to(db, command, message, user)
  if err != nil {
    common.CustomErrLog.Println(err)
    return "Unknown error", err
  }
  return fmt.Sprintf("message has been added: %v ] %v ] %v: %v", id, date, user, message), nil
}

func Remove_message(db *sql.DB, username, id string) (string, error) {
  
  _, nothingFound, err := Get_message(db, username, id) //check if message exist
  if err != nil {
    common.CustomErrLog.Println(err)
    return "Unknown error", err
  }
  
  if nothingFound {
    return "nothing was found X_X", nil
  }

  command := "DELETE FROM MESSAGE WHERE sender = $1 AND ID = $2;;" //I don't handle this to use name need to return this later XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  send_db_command_to(db, command, username, id)
  return fmt.Sprintf("message ID '%s' has been deleted successfully.", id), nil
  
}

func Get_all_messages(db *sql.DB, username string) ([]string, error) {
  command := "SELECT * FROM MESSAGE WHERE sender = $1;" //look at the spceial handling ''
  result, err := send_db_query_to(db, command, username)
  if err != nil {
    common.CustomErrLog.Println(err)
    return result, err
  }

  if len(result) <= 0 {
    return []string{"nothing was found X_X"}, nil
  }

  return result, nil
}

func Get_message(db *sql.DB, username string, id string) ([]string, bool, error) {
  command := "SELECT * FROM MESSAGE WHERE sender = $1 AND ID = $2;" //look at the spceial handling ''
  result, err := send_db_query_to(db, command, username, id)
  if err != nil {
    common.CustomErrLog.Println(err)
    return []string{"unknown error"}, false ,err
  }
  
  if len(result) <= 0 {
    return []string{"nothing was found X_X"}, true, nil
  }

  return result, false, nil
}




