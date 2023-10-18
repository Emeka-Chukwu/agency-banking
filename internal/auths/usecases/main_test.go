package usecases_handler

import (
	repo "agency-banking/internal/auths/repository/postgres"
	"agency-banking/pkg/token"
	"agency-banking/util"
	"testing"
	"time"
)

func newTestUsecase(t *testing.T, store repo.AuthRepository) Authuscase {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil
	}
	usecase := NewAuthusecase(tokenMaker, config, store)
	return usecase
}
