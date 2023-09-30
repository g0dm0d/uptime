package model

import (
	"time"

	"github.com/g0dm0d/uptime/internal/store"
)

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
}

func NewUser(u store.User) User {
	return User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}

func NewUsers(u []store.User) []User {
	var users []User
	for i := range u {
		users = append(users, NewUser(u[i]))
	}
	return users
}
