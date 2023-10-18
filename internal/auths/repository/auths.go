package auths

import (
	domain "agency-banking/domain/auths"
	"database/sql"
)

type AuthRepository interface {
	Register(data domain.Login) (domain.LoginResp, error)
	Login(data domain.Login) (domain.LoginResp, error)
}

type authRepository struct {
	DB *sql.DB
}
