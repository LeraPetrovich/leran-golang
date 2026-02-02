package storage

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (s *storage) GetCoAuthorOfAlbum(ctx context.Context, albumId int) (*core.AuthorsList, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	rows, err := s.postgres.Query(contextWait, `
	SELECT a.id, a.name, a.username
	FROM authors a
	JOIN author_album aa ON a.id = aa.author_id
	WHERE aa.album_id = $1
	`, albumId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var authors []*core.Author

	for rows.Next() {
		var author core.Author
		err := rows.Scan(&author.Id, &author.Name, &author.Username)
		if err != nil {
			return nil, err
		}

		authors = append(authors, &author)
	}

	return &core.AuthorsList{Items: authors}, nil
}

func (s *storage) GetCoAuthorOfTrack(ctx context.Context, trackId int) (*core.AuthorsList, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	rows, err := s.postgres.Query(contextWait, `
	SELECT a.id, a.name, a.username
	FROM authors a
	JOIN author_track at ON a.id = at.author_id
	WHERE at.track_id = $1
	`, trackId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var authors []*core.Author

	for rows.Next() {
		var author core.Author
		err := rows.Scan(&author.Id, &author.Name, &author.Username)
		if err != nil {
			return nil, err
		}

		authors = append(authors, &author)
	}

	return &core.AuthorsList{Items: authors}, nil
}

func (s *storage) AddCoAuthorToAlbum(ctx context.Context, idAlbum int, idAuthor int) error {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	_, err := s.postgres.Exec(contextWait, `
	INSERT INTO author_album (author_id, album_id)
	VALUES ($1, $2)
	`, idAuthor, idAlbum)

	return err
}

func (s *storage) AddCoAuthorToTrack(ctx context.Context, idTrack int, idAuthor int) error {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	_, err := s.postgres.Exec(contextWait, `
	INSERT INTO author_track (author_id, track_id)
	VALUES ($1, $2)
	`, idAuthor, idTrack)

	return err
}

func (s *storage) RemoveCoAuthorFromAlbum(ctx context.Context, idAlbum int, idAuthor int) error {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	_, err := s.postgres.Exec(contextWait, `
	DELETE FROM author_album
	WHERE album_id = $1 AND author_id = $2
	`, idAlbum, idAuthor)

	return err
}

func (s *storage) RemoveCoAuthorFromTrack(ctx context.Context, idTrack int, idAuthor int) error {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	_, err := s.postgres.Exec(contextWait, `
	DELETE FROM author_track
	WHERE track_id = $1 AND author_id = $2
	`, idTrack, idAuthor)

	return err
}
