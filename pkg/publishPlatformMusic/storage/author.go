package storage

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (s *storage) GetAuthorByUsername(ctx context.Context, username string) (*core.Author, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var author core.Author
	err := s.postgres.QueryRow(contextWait, `
	SELECT id, username, name
	FROM authors
	WHERE username = $1
	`, username).Scan(&author.Id, &author.Username, &author.Name)

	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (s *storage) GetAllAuthors(ctx context.Context) (*core.AuthorsList, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var authors []*core.Author

	rows, err := s.postgres.Query(contextWait, `
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

func (s *storage) GetAuthorFromId(ctx context.Context, authorId int) (*core.Author, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var author core.Author
	err := s.postgres.QueryRow(contextWait, `
	SELECT id, username, name
	FROM authors
	WHERE id = $1
	`, authorId).Scan(&author.Id, &author.Username, &author.Name)

	if err != nil {
		return nil, err
	}
	return &author, nil
}
