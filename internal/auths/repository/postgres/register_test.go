package auths

import (
	domain "agency-banking/domain/auths"
	"agency-banking/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) domain.User {
	arg := domain.User{
		Name:       util.RandomUsername(),
		Email:      util.RandomEmail(),
		Password:   util.RandomPassword(),
		ContactIfo: util.RandomString(45),
	}
	user, err := testQueries.Register(arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Name, arg.Name)
	require.Equal(t, user.Email, arg.Email)
	require.NotZero(t, user.CreatedAt, user.UpdatedAt)
	return user
}

func TestRegisterUser(t *testing.T) {
	createRandomUser(t)
}
