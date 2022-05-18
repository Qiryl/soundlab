package dto

import userService "github.com/grigagod/strive/proto/user"

// LoginPts are parameters for user login.
type LoginPts struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8 lte=32"`
}

func LoginPtsFromProto(req *userService.LoginReq) LoginPts {
	pts := LoginPts{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	return pts
}
