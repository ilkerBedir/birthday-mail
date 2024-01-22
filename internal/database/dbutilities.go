package database

import (
	"database/sql"
	"log"
	"project/birthday-mail/internal/config"

	_ "github.com/lib/pq"
)

var database *sql.DB
var err error

func init() {
	log.Println("DbOpen! ", config.GetConfig().Database.URL)
	database, err = sql.Open("postgres", config.GetConfig().Database.URL)
	if err != nil {
		log.Fatalln(err)
	}
}
func Close() error {
	return database.Close()
}
func SelectUsersByDate(date string) ([]User, error) {
	rows, err := database.Query("select * from users where birth_date = $1", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	var user User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name_surname, &user.Birth_date, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
func InsertUsers(users []User) {
	for _, user := range users {
		insertUser(user)
	}
}

func insertUser(user User) {
	_, err := database.Exec("insert into users (name_surname,birth_date,email) values ($1,$2,$3)", user.Name_surname, user.Birth_date, user.Email)
	if err != nil {
		log.Fatalln(err)
	}
}
