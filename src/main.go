package main

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

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	// Example usage of send_db_command_to function
	
  //create table if does not exist
  send_db_command_to(db, "CREATE TABLE IF NOT EXISTS LONG_MESSAGES (id SERIAL PRIMARY KEY, message VARCHAR(255), sender VARCHAR(50), date VARCHAR(50))")
  fmt.Println(command)
  
  add_message(db, "tal", "hello this is tal")
  remove_message(db, 5)
}

func send_db_command_to(db *sql.DB, command string) {
	_, err := db.Exec(command)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Command worked successfully!")
}

func add_message(db *sql.DB, user string, message string) {
  *&command = "INSERT INTO LONG_MESSAGES (message, sender, date) VALUES ('"+message+"', '"+user+"', '"+current_date_for_message()+"');" //look at the spceial handling ''
  send_db_command_to(db, command)
}

func remove_message(db *sql.DB, message_id int) {
  *&command = "DELETE FROM LONG_MESSAGES WHERE ID = '"+strconv.Itoa(message_id)+"';" //look at the spceial handling ''
  send_db_command_to(db, command)
}
