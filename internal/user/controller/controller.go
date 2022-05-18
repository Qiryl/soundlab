package controller

import (
	"context"

	"github.com/grigagod/strive/internal/user/dto"
	"github.com/grigagod/strive/internal/user/entity"
)

type UseCase interface {
	Register(ctx context.Context, pts dto.RegisterPts) (*entity.User, error)
	Login(ctx context.Context, pts dto.LoginPts) (*entity.User, error)
}
