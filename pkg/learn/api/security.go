package api

import (
	"context"
	"fmt"

	"example.com/learn/pkg/learn/api/oas"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserContextKey contextKey = "user"

type SecurityHandler struct {
	jwtSecret []byte
}

var _ oas.SecurityHandler = (*SecurityHandler)(nil)

func NewSecurityHandler(jwtSecret []byte) *SecurityHandler {
	return &SecurityHandler{
		jwtSecret: jwtSecret,
	}
}

func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName oas.OperationName, t oas.BearerAuth) (context.Context, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(t.Token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	ctx = context.WithValue(ctx, UserContextKey, claims)
	return ctx, nil
}

func GetUserFromContext(ctx context.Context) *Claims {
	if claims, ok := ctx.Value(UserContextKey).(*Claims); ok {
		return claims
	}
	return nil
}
