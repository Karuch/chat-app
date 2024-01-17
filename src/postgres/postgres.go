package postgres

import (
	"database/sql"
	"fmt"
	"log"
  "time"
  "strconv"
	_ "github.com/lib/pq"
)

/*  pg functions to implement in main for later use
  db := postgres.Connect_to_db()
  postgres.Add_message(db, "maor", "hello this is tal")
  postgres.Remove_message(db, 114)
  postgres.Get_all_messages(db, "' OR 1=1; --")
  postgres.Get_message(db, 4)
  defer db.Close() 
*/


const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "1598"
	dbname   = "postgres"
)

func Connect_to_db() *sql.DB {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  send_db_command_to(db, "CREATE TABLE IF NOT EXISTS LONG_MESSAGES (id SERIAL PRIMARY KEY, message VARCHAR(255), sender VARCHAR(50), date VARCHAR(50))")
  //create table if does not exist ^
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println("Successfully connected!")
  return db
}


func current_date_for_message() string {
	currentTime := time.Now()
	return strconv.Itoa(currentTime.Year()) + "/" +
		strconv.Itoa(int(currentTime.Month())) + "/" +
		strconv.Itoa(currentTime.Day()) + " " +
		strconv.Itoa(currentTime.Hour()) + ":" +
		strconv.Itoa(currentTime.Minute()) + ":" + // Change from Hour() to Minute()
		strconv.Itoa(currentTime.Second())
}

var command string 

func send_db_command_to(db *sql.DB, command string, args ...interface{}) {
	_, err := db.Exec(command, args...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Command worked successfully!")
}

var (
  id int
  sender string
  message string
  date string
)

func send_db_query_to(db *sql.DB, command string, args ...interface{}) {
  rows, err := db.Query(command, args...)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &message, &sender, &date)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(id, "]", date, "]", sender, ":", message)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}


func Add_message(db *sql.DB, user string, message string) {
  *&command = "INSERT INTO LONG_MESSAGES (message, sender, date) VALUES ($1, $2, '"+current_date_for_message()+"');" //look at the spceial handling ''
  send_db_command_to(db, command, message, user)
}

func Remove_message(db *sql.DB, message_id int) {
  *&command = "DELETE FROM LONG_MESSAGES WHERE ID = $1;" //look at the spceial handling ''
  send_db_command_to(db, command, message_id)
}

func Get_all_messages(db *sql.DB, user string) {
  *&command = "SELECT * FROM LONG_MESSAGES WHERE sender = $1;" //look at the spceial handling ''
  send_db_query_to(db, command, user)
}

func Get_message(db *sql.DB, message_id int) {
  *&command = "SELECT * FROM LONG_MESSAGES WHERE ID = $1;" //look at the spceial handling ''
  send_db_query_to(db, command, message_id)
}


