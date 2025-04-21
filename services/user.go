package services

import (
	"context"
	"session20-gin-app/models"
	"session20-gin-app/repository"
)

type UserService struct {
	repo repository.Repo
}

func (u *UserService) CreateUser(user models.User) error {
	_, err := u.repo.CreateData(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
