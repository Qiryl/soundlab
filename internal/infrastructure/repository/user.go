package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/grigagod/strive/internal/user/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	*txPool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		txPool: &txPool{Pool: pool},
	}
}

func (r *UserRepo) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	const query = `
    INSERT INTO users ("user_id", "email", "username", "full_name", "password", "bio")
    VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;
    `
	rows, err := r.txPool.conn(ctx).Query(ctx, query, user.ID, user.Email, user.Name, user.FullName, user.Password, user.Bio)
	if err == nil {
		defer rows.Close()
		err = pgxscan.ScanOne(user, rows)
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return nil, err
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}
	return user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	const query = ` SELECT * FROM users WHERE email = $1`
	var user entity.User

	rows, err := r.txPool.conn(ctx).Query(ctx, query, email)
	if err == nil {
		defer rows.Close()
		err = pgxscan.ScanOne(&user, rows)
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return nil, err
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("GetByEmail: %w", err)
	}
	return &user, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*entity.User, error) {
	const query = ` SELECT * FROM users WHERE id = $1`
	var user *entity.User

	rows, err := r.txPool.conn(ctx).Query(ctx, query, id)
	if err == nil {
		defer rows.Close()
		err = pgxscan.ScanOne(user, rows)
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return nil, err
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("GetByID: %w", err)
	}
	return user, nil
}
