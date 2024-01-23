package mail

import (
	"log"
	"net/smtp"
	"os"
	"project/birthday-mail/internal/config"
	"project/birthday-mail/internal/database"

	"github.com/jordan-wright/email"
)

var auth smtp.Auth

func init() {
	//"qozx xhvg hhab yumj"
	auth = smtp.PlainAuth("", config.GetConfig().Mail.User, config.GetConfig().MAIL_PASSWORD, "smtp.gmail.com")
}
func SendMail(user database.User) {
	to := []string{user.Email}
	e := email.NewEmail()
	e.From = config.GetConfig().Mail.User
	e.To = to
	e.Subject = "Birthday Mail"
	//e.Text = []byte("Text Body is, of course, supported!")
	//e.Text = []byte("Happy Birthday " + user.Name_surname)
	e.AttachFile("birthday.html")
	err := e.Send("smtp.gmail.com:587", auth)
	if err != nil {
		log.Fatalln(err)
	}
}

func prepareHTML(username string) ([]byte, error) {
	byteHtml, err := os.ReadFile("birthday.html")
	if err != nil {
		return nil, err
	}
	return byteHtml, nil
}
