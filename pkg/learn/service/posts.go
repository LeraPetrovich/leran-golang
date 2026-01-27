package service

import (
	"context"

	"example.com/learn/pkg/learn/core"
)

func (s *Service) GetPostList(ctx context.Context) (*core.PostList, error) {
	return s.privateAppStorage.GetPostList(ctx)
}

func (s *Service) GetPost(ctx context.Context, postId int) (*core.Post, error) {
	return s.privateAppStorage.GetPost(ctx, postId)
}

func (s *Service) CreatePost(ctx context.Context, title, description, content string) (*core.Post, error) {
	post := core.Post{
		Title:       title,
		Description: description,
		Content:     content,
	}
	return s.privateAppStorage.CreatePost(ctx, post)
}

func (s *Service) UpdatePost(ctx context.Context, postID int, title, description, content string) (*core.Post, error) {
	post := core.Post{
		Title:       title,
		Description: description,
		Content:     content,
	}
	return s.privateAppStorage.UpdatePost(ctx, postID, post)
}

func (s *Service) DeletePost(ctx context.Context, postID int) error {
	return s.privateAppStorage.DeletePost(ctx, postID)
}
