package postgres

import (
	"database/sql"
	"fmt"
	"log"
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
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  send_db_command_to(db, "CREATE TABLE IF NOT EXISTS MESSAGE (id VARCHAR(50), message VARCHAR(255), sender VARCHAR(50), date VARCHAR(50))")
  //create table if does not exist ^
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println("Successfully connected!")
  return db
}

var command string 

func send_db_command_to(db *sql.DB, command string, args ...interface{}) {
	_, err := db.Exec(command, args...)
	if err != nil {
		log.Fatal(err)
	}
}

var (
  id string
  sender string
  message string
  date string
)

func send_db_query_to(db *sql.DB, command string, args ...interface{}) []string {
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


func Add_message(db *sql.DB, user string, message string) string {
  id = common.Random_uuid()
  date = common.Current_date_for_message()
  *&command = "INSERT INTO MESSAGE (id, message, sender, date) VALUES ('"+id+"', $1, $2, '"+date+"');" //look at the spceial handling ''
  send_db_command_to(db, command, message, user)
  return fmt.Sprintf("message has been added: %v ] %v ] %v: %v", id, date, sender, message)
}

func Remove_message(db *sql.DB, id string) string {
  if Get_message(db, id) == "nothing was found X_X" { //check if message exist
    return "nothing was found X_X"
  }

  *&command = "DELETE FROM MESSAGE WHERE ID = $1;" //look at the spceial handling ''
  send_db_command_to(db, command, id)
  return fmt.Sprintf("message has been removed: '%v'", id)
  
}

func Get_all_messages(db *sql.DB, user string) []string {
  *&command = "SELECT * FROM MESSAGE WHERE sender = $1;" //look at the spceial handling ''
  return send_db_query_to(db, command, user)
}

func Get_message(db *sql.DB, id string) string{
  *&command = "SELECT * FROM MESSAGE WHERE ID = $1;" //look at the spceial handling ''
  return send_db_query_to(db, command, id)[0]
}


