package main

import (
	"fmt"
	"main/postgres"
	//"main/redis"
	"main/common"
  "os"
)

func main() {
  common.ENVinit()
  fmt.Println(common.Convert_to_int(os.Getenv("POSTGRES_PORT")))
  db := postgres.Client_connect()
  fmt.Println(postgres.Add_message(db, "maor", "hello this is tal"))
  fmt.Println(postgres.Remove_message(db, "125c1e50-6785-45f7-b40c-dda97e0130f8"))
  fmt.Println(postgres.Get_all_messages(db, "maor"))
  fmt.Println(postgres.Get_message(db, "125c1e50-6785-45f7-b40c-dda97e0130f8"))
  defer db.Close() 

  
}

