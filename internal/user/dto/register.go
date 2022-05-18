package dto

import (
	"github.com/grigagod/strive/internal/user/entity"
	userService "github.com/grigagod/strive/proto/user"
)

// RegisterPts are parameters for user registration.
type RegisterPts struct {
	Email    string  `json:"email" validate:"required,email"`
	Name     string  `json:"name" validate:"required,gte=3"`
	FullName string  `json:"full_name" validate:"required,gte=4"`
	Password string  `json:"password" validate:"required,gte=8 lte=32"`
	Bio      *string `json:"bio,omitempty"`
}

func (pts *RegisterPts) ToEntity() *entity.User {
	return &entity.User{
		Email:    pts.Email,
		Name:     pts.Name,
		FullName: pts.FullName,
		Bio:      pts.Bio,
	}
}

func RegisterPtsFromProto(req *userService.RegisterReq) RegisterPts {
	bio := req.GetBio()
	pts := RegisterPts{
		Email:    req.GetEmail(),
		Name:     req.GetName(),
		FullName: req.GetFullName(),
		Password: req.GetPassword(),
		Bio:      &bio,
	}
	return pts
}
