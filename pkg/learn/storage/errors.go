package storage

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidPostId = errors.New("invalid post id")
)
