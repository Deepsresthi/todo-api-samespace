package models

import (
	"todo-api/config"

	"github.com/gocql/gocql"
)

type User struct {
	UserID   gocql.UUID `json:"user_id"`
	Email    string     `json:"email"`
	FullName string     `json:"full_name"`
}

func (user *User) Save() error {
	user.UserID = gocql.TimeUUID()

	return config.Session.Query(`INSERT INTO users (user_id, email, full_name) VALUES (?, ?, ?)`,
		user.UserID.String(), user.Email, user.FullName).Exec()
}

func GetUserByID(userID gocql.UUID) (*User, error) {
	var user User
	if err := config.Session.Query(`SELECT user_id, email, full_name FROM users WHERE user_id = ?`,
		userID).Consistency(gocql.One).Scan(&user.UserID, &user.Email, &user.FullName); err != nil {
		return nil, err
	}
	return &user, nil
}