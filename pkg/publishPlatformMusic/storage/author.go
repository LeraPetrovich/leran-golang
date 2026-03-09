package storage

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (s *storage) GetAuthorByUsername(ctx context.Context, username string) (*core.Author, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var author core.Author
	err := s.db.QueryRow(contextWait, `
	SELECT id, username, name
	FROM authors
	WHERE username = $1
	`, username).Scan(&author.Id, &author.Username, &author.Name)

	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (s *storage) GetRefreshToken(ctx context.Context, authorId int) (string, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()
	var refreshToken string
	err := s.db.QueryRow(contextWait, `
	SELECT refresh_token
	FROM authors
	WHERE id = $1
`, authorId).Scan(&refreshToken)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (s *storage) GetAuthorByRefreshToken(ctx context.Context, refreshToken string) (*core.Author, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var author core.Author
	err := s.db.QueryRow(contextWait, `
	SELECT id, username, name
	FROM authors
	WHERE refresh_token = $1
	`, refreshToken).Scan(&author.Id, &author.Username, &author.Name)

	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (s *storage) UpdateRefreshToken(ctx context.Context, authorId int, token string) error {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	cmdTag, err := s.db.Exec(contextWait, `
	UPDATE authors
	SET refresh_token = $1
	WHERE id = $2
	`, token, authorId)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return nil
	}

	return nil

}

func (s *storage) GetAllAuthors(ctx context.Context) (*core.AuthorsList, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var authors []*core.Author

	rows, err := s.db.Query(contextWait, `
	SELECT id, username, name
	FROM authors
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var author core.Author
		err := rows.Scan(&author.Id, &author.Username, &author.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}

	return &core.AuthorsList{Items: authors}, nil
}

func (s *storage) GetAuthorById(ctx context.Context, authorId int) (*core.Author, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var author core.Author
	err := s.db.QueryRow(contextWait, `
	SELECT id, username, name
	FROM authors
	WHERE id = $1
	`, authorId).Scan(&author.Id, &author.Username, &author.Name)

	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (s *storage) GetAuthorsByAlbum(ctx context.Context, albumId int) (*core.AuthorsList, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	rows, err := s.db.Query(contextWait, `
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

func (s *storage) GetAuthorsByTrack(ctx context.Context, trackId int) (*core.AuthorsList, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	rows, err := s.db.Query(contextWait, `
	SELECT a.id, a.name, a.username
	FROM authors a
	JOIN author_album aa ON a.id = aa.author_id
	JOIN tracks t ON t.fk_id_album = aa.album_id
	WHERE t.id = $1
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
