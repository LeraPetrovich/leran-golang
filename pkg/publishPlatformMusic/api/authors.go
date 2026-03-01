package api

import (
	"context"
	"fmt"

	"example.com/learn/pkg/publishPlatformMusic/api/oas"
)

func (h *Handler) Signin(ctx context.Context, req *oas.SigninReq) (*oas.TokenPair, error) {
	userInfo, err := h.privateAppService.GetAuthorByUserName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	newAssetsToken, err := generateToken(userInfo.Id, h.jwtSecret, 30, "Access")
	if err != nil {
		return nil, fmt.Errorf("error generate assets token: %v", err)
	}
	return &oas.TokenPair{
		AccessToken:  newAssetsToken,
		RefreshToken: userInfo.RefreshToken,
	}, nil
}

func (h *Handler) GetAllAuthors(ctx context.Context) ([]oas.Author, error) {
	list, err := h.privateAppService.GetAllAuthors(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]oas.Author, len(list.Items))
	for i, a := range list.Items {
		result[i] = oas.Author{ID: a.Id, Username: a.Username, Name: a.Name}
	}
	return result, nil
}

func (h *Handler) GetAuthor(ctx context.Context, params oas.GetAuthorParams) (*oas.Author, error) {
	a, err := h.privateAppService.GetAuthorById(ctx, params.IDAuthor)
	if err != nil {
		return nil, err
	}
	return &oas.Author{ID: a.Id, Username: a.Username, Name: a.Name}, nil
}

func (h *Handler) GetAlbumsFromAuthor(ctx context.Context, params oas.GetAlbumsFromAuthorParams) ([]oas.Album, error) {
	list, err := h.privateAppService.GetAlbumsByAuthor(ctx, params.IDAuthor)
	if err != nil {
		return nil, err
	}
	result := make([]oas.Album, len(list.Items))
	for i, a := range list.Items {
		result[i] = oas.Album{ID: a.Id, Title: a.Title, Description: a.Description}
	}
	return result, nil
}

func (h *Handler) GetTracksFromAuthor(ctx context.Context, params oas.GetTracksFromAuthorParams) ([]oas.Track, error) {
	list, err := h.privateAppService.GetAllTracksByAuthor(ctx, params.IDAuthor)
	if err != nil {
		return nil, err
	}
	result := make([]oas.Track, len(list.Items))
	for i, t := range list.Items {
		result[i] = oas.Track{ID: t.Id, Name: t.Name, IDAlbum: t.IdAlbum}
	}
	return result, nil
}
