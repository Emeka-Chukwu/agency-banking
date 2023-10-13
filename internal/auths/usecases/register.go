package usecases_handler

import domain "agency-banking/domain/auths"

// @Summary Register User
// @Description register and assign a valid token to user
// @ID get-user-by-id
// @Produce json
// // @Param id path int true "User ID"
// @Success 201 {object} domain.LoginResp
// @Router /auths/register [post]
func (*authusace) Register(domain.Login) domain.LoginResp {
	panic("unimplemented")
}
