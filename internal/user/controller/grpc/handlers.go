package grpc

import (
	"context"

	"github.com/grigagod/strive/internal/user/dto"
	userService "github.com/grigagod/strive/proto/user"
)

func (c *ctrl) Register(ctx context.Context, req *userService.RegisterReq) (*userService.AuthResp, error) {
	user, err := c.uc.Register(ctx, dto.RegisterPtsFromProto(req))
	if err != nil {
		return nil, err
	}
	return userToAuthResp(user), nil
}

func (c *ctrl) Login(ctx context.Context, req *userService.LoginReq) (*userService.AuthResp, error) {
	user, err := c.uc.Login(ctx, dto.LoginPtsFromProto(req))
	if err != nil {
		return nil, err
	}
	return userToAuthResp(user), nil
}
