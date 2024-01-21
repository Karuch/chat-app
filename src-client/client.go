package main

import (
	"fmt"
	"os"
 )
 
//create main function to execute the program
func main() {
	var path string = "/home/tal/Desktop/go-chat-app/src-client/"
	create_file("access.txt", "tokenaccess", path)
	read_file("access.txt", path)
}


func create_file(fileName string, fileContent string, path string){ //will override content and create file if not exist automatically
	os.MkdirAll(path+"auth", 0777)
	fmt.Println(path+fileName)
	err := os.WriteFile(path+"auth/"+fileName, []byte(fileContent), 0666) //create a new file
	if err != nil {
	   fmt.Println(err)
	   return
	}
	fmt.Println("File is created successfully.") //print the success on the console
}

func read_file(fileName string, path string){
	content, err := os.ReadFile(path+"auth/"+fileName)

	if err != nil {
		 fmt.Println(err)
	}
   fmt.Println(string(content))
}

