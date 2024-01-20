package main

import (
	"main/common"
	"encoding/json"
    "fmt"

	

	
	//"main/postgres"
)

type Bird struct {
	Specie string
	Descriptio string
}

func main() {
	common.ENVinit()
	birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Println(birds[1].Descriptio)
	//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
	




}


/*defer func() {
	if r := recover(); r != nil {
		common.CustomErrLog.Println("Recovered from PANIC", r)
	}
}()*/