package usecases_handler

import (
	domain "agency-banking/domain/auths"
	"agency-banking/util"
)

// @Summary Login User
// @Description assign a valid token to user
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.LoginResp
// @Router /auths/login [post]
func (auth *authusace) Login(req domain.Login) (domain.LoginResp, error) {
	user, err := auth.Store.GetUser(req.Email)
	if err != nil {
		return domain.LoginResp{}, err
	}
	if err := util.CheckPassword(req.Password, user.Password); err != nil {
		return domain.LoginResp{}, err
	}
	token, payload, err := auth.TokenMaker.CreateToken(user.ID, auth.Config.AccessTokenDuration)
	if err != nil {
		return domain.LoginResp{}, err
	}
	return domain.LoginResp{
		Token:     token,
		User:      user,
		ExpiredAt: payload.ExpiredAt,
	}, err
}
