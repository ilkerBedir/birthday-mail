package database

import (
	"testing"
)

func TestInsertUsers(t *testing.T) {
	var users []User
	var user User
	user.Birth_date = "2000-01-23"
	user.Name_surname = "Mbawe Nkouudo"
	user.Email = "ilkerbedir98@gmail.com"
	users = append(users, user)
	InsertUsers(users)
}
