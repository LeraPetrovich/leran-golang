package core

import "context"

type User struct {
	ID       int
	Username string
	Password string
}

type Post struct {
	ID          int
	Title       string
	Description string
	Content     string
}

type PostList struct {
	Items []*Post
}

type AppStorage interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)

	GetPostList(ctx context.Context) (*PostList, error)
	GetPost(ctx context.Context, postId int) (*Post, error)

	CreatePost(ctx context.Context, post Post) (*Post, error)
	UpdatePost(ctx context.Context, postId int, post Post) (*Post, error)
	DeletePost(ctx context.Context, postId int) error
}
