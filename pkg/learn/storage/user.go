package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"example.com/learn/pkg/learn/core"
)

func (s *storage) GetUserByUsername(ctx context.Context, username string) (*core.User, error) {
	ctx, cancel := context.WithTimeout(ctx, fastQueryTimeout)

	defer cancel()

	var user core.User
	err := s.postgres.QueryRow(ctx, `
   SELECT id, username, password
   FROM users
   WHERE username = $1
   `, username).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
