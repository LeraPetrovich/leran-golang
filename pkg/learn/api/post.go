package api

import (
	"context"

	"example.com/learn/pkg/learn/api/oas"
)

func (h *Handler) GetPosts(ctx context.Context) ([]oas.Post, error) {
	postList, err := h.appPrivateService.GetPostList(ctx)
	if err != nil {
		return nil, err
	}

	posts := make([]oas.Post, 0, len(postList.Items))
	for _, p := range postList.Items {
		posts = append(posts, oas.Post{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			Content:     p.Content,
		})
	}
	return posts, nil
}

func (h *Handler) GetPost(ctx context.Context, params oas.GetPostParams) (*oas.Post, error) {
	post, err := h.appPrivateService.GetPost(ctx, params.PostID)
	if err != nil {
		return nil, err
	}

	return &oas.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
	}, nil
}

func (h *Handler) CreatePost(ctx context.Context, req *oas.CreatePostReq) (*oas.Post, error) {
	post, err := h.appPrivateService.CreatePost(ctx, req.Title, req.Description, req.Content)
	if err != nil {
		return nil, err
	}

	return &oas.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
	}, nil
}

func (h *Handler) UpdatePost(ctx context.Context, req *oas.UpdatePostReq, params oas.UpdatePostParams) (*oas.Post, error) {
	post, err := h.appPrivateService.UpdatePost(ctx, params.PostID, req.Title, req.Description, req.Content)
	if err != nil {
		return nil, err
	}

	return &oas.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
	}, nil
}

func (h *Handler) DeletePost(ctx context.Context, params oas.DeletePostParams) (*oas.DeletePostOK, error) {
	err := h.appPrivateService.DeletePost(ctx, params.PostID)
	if err != nil {
		return nil, err
	}

	return &oas.DeletePostOK{
		Success: oas.NewOptBool(true),
	}, nil
}
