package service

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

type ServiceInterface interface {
	GetAllAlbums(ctx context.Context) (*core.AlbumsList, error)
	GetAlbumById(ctx context.Context, albumId int) (*core.Album, error)
	CreateNewAlbum(ctx context.Context, album *core.Album, authorId int) (*core.Album, error)
	UpdateAlbum(ctx context.Context, idAlbum int, album *core.Album) (*core.Album, error)
	RemoveAlbum(ctx context.Context, idAlbum int) error
	GetTracksByAlbum(ctx context.Context, albumId int) (*core.TrackList, error)
	GetAuthorsByAlbum(ctx context.Context, albumId int) (*core.AuthorsList, error)
	UpdateAuthorAlbum(ctx context.Context, idAuthor int, idAlbum int) error
	RemoveAuthorAlbum(ctx context.Context, idAuthor int, idAlbum int) error
	GetAuthorByUserName(ctx context.Context, usename string) (*core.Author, error)
	GetRefreshToken(ctx context.Context, authorId int) (string, error)
	UpdateRefreshToken(ctx context.Context, authorId int, token string) error
	GetAllAuthors(ctx context.Context) (*core.AuthorsList, error)
	GetAuthorById(ctx context.Context, authorId int) (*core.Author, error)
	GetAlbumsByAuthor(ctx context.Context, authorId int) (*core.AlbumsList, error)
	GetAuthorByUsername(ctx context.Context, username string) (*core.Author, error)
	GetAllTracks(ctx context.Context) (*core.TrackList, error)
	GetTrackById(ctx context.Context, trackId int) (*core.Track, error)
	CreateNewTrack(ctx context.Context, track *core.Track, authorId int) (*core.Track, error)
	UpdateTrack(ctx context.Context, idTrack int, track *core.Track) (*core.Track, error)
	RemoveTrack(ctx context.Context, idTrack int) error
	GetAllTracksByAuthor(ctx context.Context, authorId int) (*core.TrackList, error)
	GetAlbumByTrack(ctx context.Context, trackId int) (*core.Album, error)
	GetAuthorsByTrack(ctx context.Context, trackId int) (*core.AuthorsList, error)
}
