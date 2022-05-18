package usecase

import (
	"context"
	"time"

	"github.com/grigagod/strive/internal/user/entity"
)

type TxRepository interface {
	TransactionContext(context.Context) (context.Context, error)
	Commit(context.Context) error
	RollBack(context.Context) error
	WithAcquire(context.Context) (context.Context, error)
	Release(context.Context)
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}

type Repository interface {
	TxRepository
	UserRepository
}

type Storage interface {
	UploadAvatar(ctx context.Context) error
	DownloadAvatar(ctx context.Context) error
}

type Cache interface {
	SetValue(ctx context.Context, key string, value any, exires time.Duration) error
	GetValue(ctx context.Context, key string) (any, error)
}

type UseCase struct {
	repo Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}
