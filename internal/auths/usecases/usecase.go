package usecases_handler

import (
	domain "agency-banking/domain/auths"
	repo "agency-banking/internal/auths/repository/postgres"
	"agency-banking/pkg/token"
	"agency-banking/util"
)

type Authuscase interface {
	Login(data domain.Login) (domain.LoginResp, error)
	Register(data domain.User) (domain.LoginResp, error)
}

type authusace struct {
	TokenMaker token.Maker
	Config     util.Config
	Store      repo.AuthRepository
}

func NewAuthusecase(token token.Maker, config util.Config, auth repo.AuthRepository) Authuscase {
	return &authusace{Config: config, TokenMaker: token, Store: auth}
}

// type ResponseData struct {
// 	Message string `json:"message"`
// 	Data    any    `json:"data,omitempty"`
// 	Error   any    `json:"error,omitempty"`
// }
