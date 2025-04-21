package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"session20-gin-app/models"
	"session20-gin-app/repository"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()

	user := models.User{
		FirstName: "Shubham",
		SeconName: "Dixit",
		Username:  "shubhamdixit",
		Password:  "securepassword",
	}
	assert.Nil(t, err)
	// Expect the INSERT exec call
	mock.ExpectExec("INSERT INTO Users").
		WithArgs(user.FirstName, user.SeconName, user.Username, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlxDB := sqlx.NewDb(db, "mysql")

	repo := repository.NewMysql(sqlxDB)
	us := UserService{repo: repo}
	err = us.CreateUser(user)
	assert.Nil(t, err)
}
