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
	//database.GetTip("Use strong and unique passwords for each of your online accounts, and consider using a password manager to generate and store them securely.")
	//database.GetTip("Keep your software and operating system up to date with the latest security patches, as these often address known vulnerabilities.")
	//database.GetTip("Be cautious when clicking on links or downloading attachments from unknown or suspicious sources, as they may contain malware.")
	//database.GetTip("Use a reputable antivirus program to help protect your computer from malware and other security threats.")
	server.StartServer()
}
