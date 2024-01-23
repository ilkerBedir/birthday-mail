package main

import (
	"fmt"
	"log"
	"project/birthday-mail/internal/config"
	"project/birthday-mail/internal/database"
	"project/birthday-mail/internal/mail"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting...")
	config.GetConfig()
	defer database.Close()
	now := time.Now()
	now2 := fmt.Sprintf("%s", now.Format("2006-01-02"))
	fmt.Println(now2)
	users, err := database.SelectUsersByDate(now2)
	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range users {
		mail.SendMail(user)
	}
}
