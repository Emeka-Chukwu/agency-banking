package auths

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user2.Name, user1.Name)
	require.Equal(t, user2.Email, user1.Email)
	require.NotZero(t, user2.CreatedAt, user1.CreatedAt)
	require.NotZero(t, user2.UpdatedAt, user1.UpdatedAt)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second*1)
}
