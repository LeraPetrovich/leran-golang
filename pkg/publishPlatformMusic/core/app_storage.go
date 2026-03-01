package core

import (
	"context"
)

type Author struct {
	Id           int
	Username     string
	Name         string
	RefreshToken string
}

type AuthorsList struct {
	Items []*Author
}

type Album struct {
	Id          int
	Title       string
	Description string
}

type AlbumsList struct {
	Items []*Album
}

type Track struct {
	Id      int
	Name    string
	IdAlbum int
}

type TrackList struct {
	Items []*Track
}

type AuthorAlbum struct {
	IdAuthor int
	IdAlbum  int
}

type AppStorage interface {
	//auth
	GetAuthorByUsername(ctx context.Context, username string) (*Author, error)
	GetAuthorByRefreshToken(ctx context.Context, refreshToken string) (*Author, error)
	GetRefreshToken(ctx context.Context, authorId int) (string, error)
	UpdateRefreshToken(ctx context.Context, authorId int, token string) error

	//get co-author
	GetAuthorsByAlbum(ctx context.Context, albumId int) (*AuthorsList, error)
	GetAuthorsByTrack(ctx context.Context, trackId int) (*AuthorsList, error)

	//getAuthorInfo
	GetAuthorById(ctx context.Context, authorId int) (*Author, error)
	GetAllAuthors(ctx context.Context) (*AuthorsList, error)

	//get track
	GetTrackById(ctx context.Context, trackId int) (*Track, error)
	GetTracksByAlbum(ctx context.Context, albumId int) (*TrackList, error)
	GetAllTracksByAuthor(ctx context.Context, authorId int) (*TrackList, error)
	GetAllTracks(ctx context.Context) (*TrackList, error)

	//get album
	GetAlbumsByAuthor(ctx context.Context, authorId int) (*AlbumsList, error)
	GetAlbumById(ctx context.Context, albumId int) (*Album, error)
	GetAlbumByTrack(ctx context.Context, trackId int) (*Album, error)
	GetAllAlbums(ctx context.Context) (*AlbumsList, error)

	//create
	CreateNewAlbum(ctx context.Context, album *Album, authorId int) (*Album, error)
	CreateNewTrack(ctx context.Context, track *Track, authorId int) (*Track, error)

	//update
	UpdateAlbum(ctx context.Context, idAlbum int, album *Album) (*Album, error)
	UpdateTrack(ctx context.Context, idTrack int, track *Track) (*Track, error)
	UpdateAuthorAlbum(ctx context.Context, idAuthor int, idAlbum int) error

	//join / split
	RemoveTrack(ctx context.Context, idTrack int) error
	RemoveAlbum(ctx context.Context, idAlbum int) error
	RemoveAuthorAlbum(ctx context.Context, idAuthor int, idAlbum int) error
}
