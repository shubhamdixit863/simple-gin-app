package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"session20-gin-app/models"
	"strconv"
)

type Mysql struct {
	conn *sqlx.DB
}

func (m Mysql) CreateData(ctx context.Context, user models.User) (string, error) {
	result, err := m.conn.ExecContext(ctx, `INSERT INTO Users (FirstName, SecondName, UserName, password) VALUES (?, ?, ?, ?)`, user.FirstName, user.SeconName, user.Username, user.Password)

	if err != nil {
		log.Println(err)

	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(id)), nil
}

func NewMysql(conn *sqlx.DB) Repo {
	return Mysql{conn: conn}
}
