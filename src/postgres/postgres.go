package postgres

import (
	"database/sql"
	"fmt"
	"log"
  "time"
  "strconv"
	_ "github.com/lib/pq"
)

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

func send_db_command_to(db *sql.DB, command string) {
	_, err := db.Exec(command)
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

func send_db_query_to(db *sql.DB, command string) {
  rows, err := db.Query(command)
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&id, &message, &sender, &date) //the order will scan the rows by the order they were saved in the db so it matters
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(id,"]", date ,"]", sender,":", message)
  }
  err = rows.Err()
  if err != nil {
    log.Fatal(err)
  }
}


func Add_message(db *sql.DB, user string, message string) {
  *&command = "INSERT INTO LONG_MESSAGES (message, sender, date) VALUES ('"+message+"', '"+user+"', '"+current_date_for_message()+"');" //look at the spceial handling ''
  send_db_command_to(db, command)
}

func Remove_message(db *sql.DB, message_id int) {
  *&command = "DELETE FROM LONG_MESSAGES WHERE ID = '"+strconv.Itoa(message_id)+"';" //look at the spceial handling ''
  send_db_command_to(db, command)
}

func Get_all_messages(db *sql.DB, user string) {
  *&command = "SELECT * FROM LONG_MESSAGES WHERE sender = '"+user+"';" //look at the spceial handling ''
  send_db_query_to(db, command)
}

func Get_message(db *sql.DB, message_id int) {
  *&command = "SELECT * FROM LONG_MESSAGES WHERE ID = '"+strconv.Itoa(message_id)+"';" //look at the spceial handling ''
  send_db_query_to(db, command)
}


