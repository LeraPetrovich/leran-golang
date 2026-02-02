package core

import (
	"context"
)

type Author struct {
	Id       int
	Username string
	Name     string
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

type AuthorTrack struct {
	IdAuthor int
	IdTrack  int
}

type AppStorage interface {
	//auth
	GetAuthorByUsername(ctx context.Context, username string) (*Author, error)

	//get co-author
	GetCoAuthorOfAlbum(ctx context.Context, albumId int) (*AuthorsList, error)
	GetCoAuthorOfTrack(ctx context.Context, trackId int) (*AuthorsList, error)

	//getAuthorInfo
	GetAuthorFromId(ctx context.Context, authorId int) (*Author, error)

	//get track
	GetTrackById(ctx context.Context, trackId int) (*Track, error)
	GetTracksInAlbum(ctx context.Context, albumId int) (*TrackList, error)
	GetAllTrackFromAuthor(ctx context.Context, authorId int) (*TrackList, error)

	//get album
	GetAllAlbumsFromAuthor(ctx context.Context, authorId int) (*AlbumsList, error)
	GetAlbumFromId(ctx context.Context, albumId int) (*Album, error)
	GetAlbumFromTrackId(ctx context.Context, trackId int) (*Album, error)

	//create
	CreateNewAlbum(ctx context.Context, album *Album, authorId int) (*Album, error)
	CreateNewTrack(ctx context.Context, track *Track, authorId int) (*Track, error)

	//update
	UpdateAlbum(ctx context.Context, idAlbum int, album *Album) (*Album, error)
	UpdateTrack(ctx context.Context, idTrack int, track *Track) (*Track, error)

	//join / split
	AddCoAuthorToAlbum(ctx context.Context, idAlbum int, idAuthor int) error
	AddCoAuthorToTrack(ctx context.Context, idTrack int, idAuthor int) error

	RemoveCoAuthorFromAlbum(ctx context.Context, idAlbum int, idAuthor int) error
	RemoveCoAuthorFromTrack(ctx context.Context, idAlbum int, idAuthor int) error
	RemoveTrack(ctx context.Context, idTrack int) error
	RemoveAlbum(ctx context.Context, idAlbum int) error
}
