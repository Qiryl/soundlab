package repository

import (
	"context"
	"testing"

	"github.com/grigagod/strive/internal/user/entity"
	"github.com/hatch-studio/pgtools/sqltest"
)

func TestUserRepo(t *testing.T) {
	t.Parallel()

	migration := sqltest.New(t, sqltest.Options{
		Path: "../../../migrations/user",
	})

	pool := migration.Setup(context.Background(), mainURL)
	repo := NewUserRepo(pool)

	t.Run("Create", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			user := entity.User{
				ID:       "1",
				Email:    "test",
				Name:     "test",
				FullName: "test",
				Password: "test",
			}
			_, err := repo.Create(context.Background(), &user)
			if err != nil {
				t.Error(err)
			}
		})
	})
	t.Run("GetByEmail", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			createUsers(t, repo, []*entity.User{{ID: "2", Email: "exmp@mail"}})
			got, err := repo.GetByEmail(context.Background(), "exmp@mail")
			if err != nil || got.ID != "2" {
				t.Error(err)
			}
		})
		t.Run("not found", func(t *testing.T) {
			got, err := repo.GetByEmail(context.Background(), "mail@exmp")
			if err != nil || got != nil {
				t.Error(err)
			}
		})
	})
}
