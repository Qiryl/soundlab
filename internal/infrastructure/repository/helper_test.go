package repository

import (
	"context"
	"testing"

	"github.com/grigagod/strive/internal/user/entity"
)

func createUsers(t testing.TB, repo *UserRepo, users []*entity.User) {
	for _, u := range users {
		if _, err := repo.Create(context.Background(), u); err != nil {
			t.Error(err)
		}
	}
}
