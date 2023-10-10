package ressponse

import "codebase.sample/model"

type userResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	} `json:"user"`
}

func NewUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	// r.User.Token = utils.GenerateJWT(u.ID)
	return r
}
