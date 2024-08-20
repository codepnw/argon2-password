package store

import (
	"fmt"

	"github.com/codepnw/argon2password/types"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) UserStore {
	return UserStore{db: db}
}

func (us *UserStore) InsertUser(user *types.User) (*types.User, error) {
	var result types.User
	query := `insert into users (email, password_hash) values ($1, $2) returning *;`

	err := us.db.Get(&result, query, user.Email, user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("InsertUser: %w", err)
	}
	return &result, err
}

func (us *UserStore) GetUserByEmail(email string) (*types.User, error) {
	var result types.User
	query := `select * from users where email = $1;`

	err := us.db.Get(&result, query, email)
	if err != nil {
		return nil, fmt.Errorf("GetUserByEmail: %w", err)
	}
	return &result, err
}