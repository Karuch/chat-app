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
  postgres.Add_message(db, "maor", "hello this is tal")
  //postgres.Remove_message(db, "751266b8-10cb-48d7-8c6d-0734c608614c")
  fmt.Println(postgres.Get_all_messages(db, "maor"))
  //fmt.Println(postgres.Get_message(db, "0e99ae35-b07c-4498-8816-e6e3da16563e"))
  defer db.Close() 

  
}

