package api

import (
	"context"
	"time"

	"example.com/learn/pkg/learn/api/oas"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (h *Handler) Signin(ctx context.Context, req *oas.SigninReq) (*oas.SigninOK, error) {
	err := h.appPrivateService.SignIn(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	expiration := time.Now().Add(600 * time.Minute) //toke life
	claims := Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(h.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &oas.SigninOK{
		Token: token,
	}, nil
}
