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
  postgres.Get_all_messages(postgres.Connect_to_db(), "maor")
  postgres.Add_message(postgres.Connect_to_db(), "maor", "olaaa maorrr")
}

