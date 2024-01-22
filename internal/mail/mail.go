package mail

import (
	"log"
	"net/smtp"
	"os"
	"project/birthday-mail/internal/config"
	"project/birthday-mail/internal/database"
	"strings"
)

var auth smtp.Auth

func init() {
	//"qozx xhvg hhab yumj"
	auth = smtp.PlainAuth("", config.GetConfig().Mail.User, config.GetConfig().MAIL_PASSWORD, "smtp.gmail.com")
}
func SendMail(user database.User) {
	to := []string{user.Email}
	html, err := prepareHTML(user.Name_surname)
	if err != nil {
		log.Fatalln(err)
	}
	msg := []byte(
		"Subject: Why aren’t you using Mailtrap yet?\r\n" +
			"Content-Type: text/html; charset=utf-8" +
			"\r\n" +
			html,
	)
	err = smtp.SendMail("smtp.gmail.com:587", auth, config.GetConfig().Mail.User, to, msg)
	if err != nil {
		log.Fatalln(err)
	}
}

func prepareHTML(username string) (string, error) {
	byteHtml, err := os.ReadFile("birthday.html")
	if err != nil {
		return "", err
	}
	stringHtml := string(byteHtml)
	stringHtml = strings.ReplaceAll(stringHtml, "Amal", username)
	return stringHtml, nil
}
