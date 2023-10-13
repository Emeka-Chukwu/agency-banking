package usecases_handler

import domain "agency-banking/domain/auths"

// @Summary Login User
// @Description assign a valid token to user
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.LoginResp
// @Router /auths/login [post]
func (*authusace) Login(domain.Login) domain.LoginResp {
	panic("unimplemented")
}
