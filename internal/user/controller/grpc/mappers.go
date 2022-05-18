package grpc

import (
	"github.com/grigagod/strive/internal/user/entity"
	userService "github.com/grigagod/strive/proto/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func userToAuthResp(user *entity.User) *userService.AuthResp {
	return &userService.AuthResp{
		ID:        user.ID,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
