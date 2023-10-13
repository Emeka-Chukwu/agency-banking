package usecases_handler

import (
	domain "agency-banking/domain/auths"
	"agency-banking/pkg/token"
	"agency-banking/util"
)

type Authuscase interface {
	Login(domain.Login) (domain.LoginResp, error)
	Register(domain.Login) (domain.LoginResp, error)
}

type authusace struct {
	TokenMaker token.Maker
	Config     util.Config
}

func NewAuthusecase(token token.Maker, config util.Config) Authuscase {
	return &authusace{Config: config, TokenMaker: token}
}
