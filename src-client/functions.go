package main

import (
	"fmt"
	"os"
 )
 
var path string = "/home/tal/Desktop/go-chat-app/src-client/"

//create main function to execute the program
func main() {
	
	Create_file("access.txt", "tokenaccess", path)
	Read_file("access.txt", path)
}


func Create_file(fileName string, fileContent string, path string){ //will override content and create file if not exist automatically
	os.MkdirAll(path+"auth", 0777)
	fmt.Println(path+fileName)
	err := os.WriteFile(path+"auth/"+fileName, []byte(fileContent), 0666) //create a new file
	if err != nil {
	   fmt.Println(err)
	   return
	}
	fmt.Println("File is created successfully.") //print the success on the console
}

func Read_file(fileName string, path string){
	content, err := os.ReadFile(path+"auth/"+fileName)

	if err != nil {
		 fmt.Println(err)
	}
   fmt.Println(string(content))
}

func Handle_server_answer(token string, status string){
    switch status {
    case "login_is_true":
		//will get and create refresh
		Create_file("refresh.txt", token, path)
        fmt.Println("login is correct.")
    case "login_is_wrong":
		//will try login
        fmt.Println("username or password invalid. try again")
    case "access_is_true":
		//will do nothing
        fmt.Println("accesstoken true")
	case "access_is_wrong":
		//will try request
        fmt.Println("accesstoken wrong")
	case "refresh_is_true":
		//will get and create access
		Create_file("access.txt", token, path)
        fmt.Println("refreshtoken true")
	case "refresh_is_wrong":
        fmt.Println("refreshtoken wrong")
		//will try login
	case "half_time_refresh":
		//will get and create refresh and get and create access
		fmt.Println("half_time_refresh")
    }
}

