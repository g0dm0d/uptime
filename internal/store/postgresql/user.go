package postgresql

import (
	"database/sql"
	"github.com/g0dm0d/uptime/internal/store"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) store.UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) CreateUser(opts store.CreateUserOpts) (id int, err error) {
	req := s.db.QueryRow("SELECT * FROM create_user($1, $2)",
		opts.Username, opts.Password)

	err = req.Scan(&id)

	return id, err
}

func (s *UserStore) GetUserByUsername(username string) (store.User, error) {
	var user store.User
	req := s.db.QueryRow("SELECT * FROM get_user_by_username($1)", username)
	err := req.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
