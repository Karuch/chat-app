package main

import (
  "main/postgres"
  "fmt"
)

func main() {
  db := postgres.Connect_to_db()
  postgres.Add_message(db, "maor", "hello this is tal")
  postgres.Remove_message(db, 114)
  postgres.Get_all_messages(db, "' OR 1=1; --")
  postgres.Get_message(db, 4)
  defer db.Close()
}

