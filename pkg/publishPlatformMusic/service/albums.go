package service

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (s *Service) GetAllAlbums(ctx context.Context) (*core.AlbumsList, error) {
	return s.privateAppStorage.GetAllAlbums(ctx)
}

func (s *Service) GetAlbumById(ctx context.Context, albumId int) (*core.Album, error) {
	return s.privateAppStorage.GetAlbumById(ctx, albumId)
}

func (s *Service) CreateNewAlbum(ctx context.Context, album *core.Album, authorId int) (*core.Album, error) {
	return s.privateAppStorage.CreateNewAlbum(ctx, album, authorId)
}

func (s *Service) UpdateAlbum(ctx context.Context, idAlbum int, album *core.Album) (*core.Album, error) {
	return s.privateAppStorage.UpdateAlbum(ctx, idAlbum, album)
}

func (s *Service) RemoveAlbum(ctx context.Context, idAlbum int) error {
	return s.privateAppStorage.RemoveAlbum(ctx, idAlbum)
}

func (s *Service) GetTracksByAlbum(ctx context.Context, albumId int) (*core.TrackList, error) {
	return s.privateAppStorage.GetTracksByAlbum(ctx, albumId)
}

func (s *Service) GetAuthorsByAlbum(ctx context.Context, albumId int) (*core.AuthorsList, error) {
	return s.privateAppStorage.GetAuthorsByAlbum(ctx, albumId)
}

func (s *Service) UpdateAuthorAlbum(ctx context.Context, idAuthor int, idAlbum int) error {
	return s.privateAppStorage.UpdateAuthorAlbum(ctx, idAuthor, idAlbum)
}

func (s *Service) RemoveAuthorAlbum(ctx context.Context, idAuthor int, idAlbum int) error {
	return s.privateAppStorage.RemoveAuthorAlbum(ctx, idAuthor, idAlbum)
}
