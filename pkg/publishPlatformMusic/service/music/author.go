package music

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (s *Service) GetAuthorByUserName(ctx context.Context, usename string) (*core.Author, error) {
	author, err := s.privateAppStorage.GetAuthorByUsername(ctx, usename)
	if err != nil {
		return nil, err
	}
	resAuthor := core.Author{
		Id:           author.Id,
		Name:         author.Name,
		RefreshToken: author.RefreshToken,
	}
	return &resAuthor, nil
}

func (s *Service) GetRefreshToken(ctx context.Context, authorId int) (string, error) {
	token, err := s.privateAppStorage.GetRefreshToken(ctx, authorId)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Service) UpdateRefreshToken(ctx context.Context, authorId int, token string) error {
	err := s.privateAppStorage.UpdateRefreshToken(ctx, authorId, token)
	if err != nil {
		return err
	}
	return nil
}
