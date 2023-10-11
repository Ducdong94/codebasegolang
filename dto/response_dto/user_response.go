package response_dto

import (
	"codebase.sample/model"
	"codebase.sample/utils"
)

type userResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func NewUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.Username = u.Username
	r.Email = u.Email
	r.Token = utils.GenerateJWT(u.Id.Hex())
	return r
}
