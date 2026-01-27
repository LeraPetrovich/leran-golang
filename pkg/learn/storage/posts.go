package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"example.com/learn/pkg/learn/core"
)

func (s *storage) GetPostList(ctx context.Context) (*core.PostList, error) {
	ctx, cancel := context.WithTimeout(ctx, fastQueryTimeout)

	defer cancel()

	rows, err := s.postgres.Query(ctx, `
	SELECT id title description content
	FROM posts
	ORDER BY id DESC
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*core.Post
	for rows.Next() {
		var post core.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return &core.PostList{Items: posts}, nil
}

func (s *storage) GetPost(ctx context.Context, postId int) (*core.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, fastQueryTimeout)

	defer cancel()
	var post core.Post
	err := s.postgres.QueryRow(ctx, `
	SELECT id, title, description, content
	FROM posts
	WHERE id = $1
	`, postId).Scan(&post.ID, &post.Title, &post.Description, &post.Content)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrInvalidPostId
		}
		return nil, err
	}
	return &post, nil
}

func (s *storage) CreatePost(ctx context.Context, post core.Post) (*core.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var id int
	err := s.postgres.QueryRow(ctx, `
INSERT INTO posts (title, description, content)
VALUES ($1, $2, $3)
RETURNING id`, post.Title, post.Description, post.Content).Scan(&id)
	if err != nil {
		return nil, err
	}

	post.ID = id
	return &post, nil
}

func (s *storage) UpdatePost(ctx context.Context, postId int, post core.Post) (*core.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	cmdTag, err := s.postgres.Exec(ctx, `
	UPDATE posts
	SET title = $1, description = $2, content = $3
	WHERE id = $4
	`, post.Title, post.Description, post.Content, postId)

	if err != nil {
		return nil, err
	}
	if cmdTag.RowsAffected() == 0 {
		return nil, nil
	}
	post.ID = postId
	return &post, nil
}

func (s *storage) DeletePost(ctx context.Context, postId int) error {
	ctx, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	_, err := s.postgres.Exec(ctx, `
DELETE FROM posts
WHERE id = $1`, postId)
	return err
}
