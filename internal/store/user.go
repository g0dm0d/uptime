package store

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
}

type CreateUserOpts struct {
	Username string
	Password string
}

type UserStore interface {
	CreateUser(opts CreateUserOpts) (int, error)
	GetUserByUsername(username string) (User, error)
}
