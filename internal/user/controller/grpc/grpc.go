package grpc

import "github.com/grigagod/strive/internal/user/controller"

type ctrl struct {
	uc controller.UseCase
}

func NewUserControllerGRPC(uc controller.UseCase) *ctrl {
	return &ctrl{uc: uc}
}
