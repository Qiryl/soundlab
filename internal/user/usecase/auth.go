package usecase

import (
	"context"
	"fmt"

	"github.com/grigagod/strive/internal/user/dto"
	"github.com/grigagod/strive/internal/user/entity"
)

// Register implements user registration usecase.
func (uc *UseCase) Register(ctx context.Context, pts dto.RegisterPts) (*entity.User, error) {
	fn := func(ctx context.Context, pts dto.RegisterPts) (*entity.User, error) {
		_, err := uc.repo.GetByEmail(ctx, pts.Email)
		if err != nil {
			return nil, err
		}
		usr := pts.ToEntity()
		if err := usr.PrepareCreate(); err != nil {
			return nil, err
		}
		created, err := uc.repo.Create(ctx, usr)
		if err != nil {
			return nil, err
		}
		created.SanitizePassword()
		return created, nil
	}

	ctx, err := uc.repo.TransactionContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("Register: %w", err)
	}

	user, err := fn(ctx, pts)
	if err != nil {
		_ = uc.repo.RollBack(ctx)
		return nil, fmt.Errorf("Register: %w", err)
	}

	if err := uc.repo.Commit(ctx); err != nil {
		return nil, fmt.Errorf("Register: %w", err)
	}
	return user, nil
}

// Login implements user log-in usecase.
func (uc *UseCase) Login(ctx context.Context, pts dto.LoginPts) (*entity.User, error) {
	found, err := uc.repo.GetByEmail(ctx, pts.Email)
	if err != nil {
		return nil, fmt.Errorf("Login: %w", err)
	}
	if err := found.ComparePasswords(pts.Password); err != nil {
		return nil, fmt.Errorf("Login: %w", err)
	}
	found.SanitizePassword()
	return found, nil
}
