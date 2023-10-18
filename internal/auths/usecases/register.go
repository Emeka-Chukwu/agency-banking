package usecases_handler

import (
	domain "agency-banking/domain/auths"
	"agency-banking/util"
)

// @Summary Register User
// @Description register and assign a valid token to user
// @ID get-user-by-id
// @Produce json
// // @Param id path int true "User ID"
// @Success 201 {object} domain.LoginResp
// @Router /auths/register [post]
func (auth *authusace) Register(req domain.User) (domain.LoginResp, error) {
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return domain.LoginResp{}, err
	}
	req.Password = hashPassword
	user, err := auth.Register(req)
	if err != nil {
		return domain.LoginResp{}, err
	}
	token, payload, err := auth.TokenMaker.CreateToken(user.ID, auth.Config.AccessTokenDuration)

	return domain.LoginResp{
		Token:     token,
		User:      user.User,
		ExpiredAt: payload.ExpiredAt,
	}, err
}
