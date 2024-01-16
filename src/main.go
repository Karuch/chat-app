package main

import (
  "main/postgres"
)

func main() {
  db := postgres.Connect_to_db()
  
    
  postgres.Add_message(db, "tal", "hello this is tal")
  postgres.Remove_message(db, 5)
  postgres.Get_all_messages(db, "tal")
  postgres.Get_message(db, 4)
  defer db.Close()
}

