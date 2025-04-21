package repository

import (
	"context"
	"session20-gin-app/models"
)

type Repo interface {
	CreateData(ctx context.Context, user models.User) (string, error)
}
