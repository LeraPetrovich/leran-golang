package api

import (
	"context"
	"fmt"
	"time"

	"example.com/learn/pkg/publishPlatformMusic/api/oas"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserContextKey contextKey = "user"

type Claims struct {
	UserId    int
	TokenType string
	jwt.RegisteredClaims
}

type SecurityHandler struct {
	jwtSecret []byte
}

var _ oas.SecurityHandler = (*SecurityHandler)(nil)

func NewSecurityHandler(jwtSecret []byte) *SecurityHandler {
	return &SecurityHandler{
		jwtSecret: jwtSecret,
	}
}

func (sh *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName oas.OperationName, t oas.BearerAuth) (context.Context, error) {
	token, err := jwt.ParseWithClaims(t.Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return sh.jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	claims := token.Claims.(*Claims)
	if claims.TokenType != "access" {
		return nil, fmt.Errorf("invalid token type")
	}

	cxtValue := context.WithValue(ctx, UserContextKey, claims)
	return cxtValue, nil
}

func (h *Handler) Refresh(ctx context.Context, req *oas.RefreshReq) (*oas.TokenPair, error) {
	token, err := jwt.ParseWithClaims(req.RefreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return h.jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(*Claims)

	if claims.TokenType != "refresh" {
		return nil, fmt.Errorf("invalid token type")
	}

	currentToken, err := h.privateAppService.GetRefreshToken(ctx, claims.UserId)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if currentToken != req.RefreshToken {
		return nil, fmt.Errorf("unknown token")
	}

	accessToken, err := generateToken(claims.UserId, h.jwtSecret, 30, "Access")
	if err != nil {
		return nil, fmt.Errorf("error generate access token: %w", err)
	}

	refreshToken, err := generateToken(claims.UserId, h.jwtSecret, 600, "Refresh")
	if err != nil {
		return nil, fmt.Errorf("error generate refresh token: %w", err)
	}

	if err := h.privateAppService.UpdateRefreshToken(ctx, claims.UserId, refreshToken); err != nil {
		return nil, fmt.Errorf("error update token in database: %w", err)
	}

	return &oas.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateToken(userId int, secret []byte, tokenLive int, tokenType string) (string, error) {
	expiresAt := time.Now().Add(time.Duration(tokenLive) * time.Minute)
	claims := Claims{
		UserId:    userId,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)

	if err != nil {
		return "", err
	}
	return token, nil
}
