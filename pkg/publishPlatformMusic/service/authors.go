package service

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (s *Service) GetAllAuthors(ctx context.Context) (*core.AuthorsList, error) {
	return s.privateAppStorage.GetAllAuthors(ctx)
}

func (s *Service) GetAuthorById(ctx context.Context, authorId int) (*core.Author, error) {
	return s.privateAppStorage.GetAuthorById(ctx, authorId)
}

func (s *Service) GetAlbumsByAuthor(ctx context.Context, authorId int) (*core.AlbumsList, error) {
	return s.privateAppStorage.GetAlbumsByAuthor(ctx, authorId)
}

func (s *Service) GetAuthorByUsername(ctx context.Context, username string) (*core.Author, error) {
	return s.privateAppStorage.GetAuthorByUsername(ctx, username)
}
