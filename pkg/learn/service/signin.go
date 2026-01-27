package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignIn(ctx context.Context, username, password string) error {
	user, err := s.privateAppStorage.GetUserByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return fmt.Errorf("invalid credentials")
	}
	return nil
}
