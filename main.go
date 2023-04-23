package main

import (
	"be/serv/database"
	"be/serv/server"
	"fmt"
)

func main() {
	fmt.Println("main !")
	err := database.InitDB()
	if err != nil {
		fmt.Println("error connecting to the db")
		return
	}
	server.StartServer()
}
