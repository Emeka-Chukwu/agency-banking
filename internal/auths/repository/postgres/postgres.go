package auths

import (
	domain "agency-banking/domain/auths"
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 5

type AuthRepository interface {
	Register(data domain.User) (domain.User, error)
	GetUser(email string) (domain.User, error)
}

type authRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{DB: db}
}
