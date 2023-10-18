package auths

import (
	domain "agency-banking/domain/auths"
	"context"
)

// Register implements AuthRepository.
func (auth *authRepository) Register(data domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `insert into agents (name, email, password, contact_info)
		values ($1, $2, $3, $4) returning id, name, email, password, contact_info, created_at, updated_at`
	var user domain.User
	err := auth.DB.QueryRowContext(ctx, stmt,
		data.Name,
		data.Email,
		data.Password,
		data.ContactIfo,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.ContactIfo,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}
