package auths

import (
	domain "agency-banking/domain/auths"
	"context"
)

// Login implements AuthRepository.
func (auth *authRepository) GetUser(email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `select id, username, email, password_hash, created_at, updated_at from users where email=$1`
	var user domain.User
	err := auth.DB.QueryRowContext(ctx, stmt, email).
		Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	return user, err
}
