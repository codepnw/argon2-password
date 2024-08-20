package store_test

import (
	"database/sql"
	"strings"

	"github.com/codepnw/argon2password/store"
	"github.com/codepnw/argon2password/types"
)

func (s *StoreTestSuite) TestInsertUser() {
	passwordHash := "test"
	user := types.User{
		Email:        "user1@mail.com",
		PasswordHash: &passwordHash,
	}
	us := store.NewUserStore(s.db)

	actual, err := us.InsertUser(&user)
	s.NoError(err)

	s.Greater(actual.ID, 0)
	s.Equal(user.Email, actual.Email)
	s.Equal(user.PasswordHash, actual.PasswordHash)
}

func (s *StoreTestSuite) TestGetUserByEmail() {
	passwordHash := "test"
	newUser := types.User{
		Email:        "example@mail.com",
		PasswordHash: &passwordHash,
	}

	us := store.NewUserStore(s.db)
	user, err := us.InsertUser(&newUser)
	s.NoError(err)

	examples := []string{user.Email, strings.ToUpper(user.Email), "Example@Mail.Com"}
	for _, email := range examples {
		actual, err := us.GetUserByEmail(email)
		s.NoError(err)
		s.Equal(user.ID, actual.ID)
	}

	actual, err := us.GetUserByEmail("norows@mail.com")
	s.ErrorIs(err, sql.ErrNoRows)
	s.Nil(actual)
}
