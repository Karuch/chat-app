package main

import (
  "database/sql"
  "fmt"
  "log"
  _ "github.com/lib/pq"
)
//test
const (
  host     = "172.17.0.2"
  port     = 5432
  user     = "postgres"
  password = "1598"
  dbname   = "postgres"
)

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

  var dbName string
  err = db.QueryRow("SELECT current_database()").Scan(&dbName)
  if err != nil {
	  log.Fatal(err)
  }
  fmt.Println(dbName)
}