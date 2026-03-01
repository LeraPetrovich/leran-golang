package api

import (
	"context"
	"fmt"
	"net/http"

	"example.com/learn/pkg/publishPlatformMusic/api/oas"
	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (h *Handler) NewError(ctx context.Context, err error) *oas.ErrorStatusCode {
	return &oas.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response:   oas.Error{Error: err.Error()},
	}
}

func (h *Handler) GetAllAlbums(ctx context.Context) ([]oas.Album, error) {
	list, err := h.privateAppService.GetAllAlbums(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]oas.Album, len(list.Items))
	for i, a := range list.Items {
		result[i] = oas.Album{ID: a.Id, Title: a.Title, Description: a.Description}
	}
	return result, nil
}

func (h *Handler) GetAlbumFromId(ctx context.Context, params oas.GetAlbumFromIdParams) (*oas.Album, error) {
	a, err := h.privateAppService.GetAlbumById(ctx, params.IDAlbum)
	if err != nil {
		return nil, err
	}
	return &oas.Album{ID: a.Id, Title: a.Title, Description: a.Description}, nil
}

func (h *Handler) CreateNewAlbum(ctx context.Context, req *oas.CreateNewAlbumReq) (*oas.Album, error) {
	claims, ok := ctx.Value(UserContextKey).(*Claims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	album := &core.Album{Title: req.Title, Description: req.Description}
	created, err := h.privateAppService.CreateNewAlbum(ctx, album, claims.UserId)
	if err != nil {
		return nil, err
	}
	return &oas.Album{ID: created.Id, Title: created.Title, Description: created.Description}, nil
}

func (h *Handler) UpdateAlbum(ctx context.Context, req *oas.UpdateAlbumReq, params oas.UpdateAlbumParams) (*oas.Album, error) {
	album := &core.Album{Title: req.Title, Description: req.Description}
	updated, err := h.privateAppService.UpdateAlbum(ctx, params.IDAlbum, album)
	if err != nil {
		return nil, err
	}
	if updated == nil {
		return nil, fmt.Errorf("album not found")
	}
	return &oas.Album{ID: updated.Id, Title: updated.Title, Description: updated.Description}, nil
}

func (h *Handler) DeleteAlbum(ctx context.Context, params oas.DeleteAlbumParams) error {
	return h.privateAppService.RemoveAlbum(ctx, params.IDAlbum)
}

func (h *Handler) GetTracksFromAlbum(ctx context.Context, params oas.GetTracksFromAlbumParams) ([]oas.Track, error) {
	list, err := h.privateAppService.GetTracksByAlbum(ctx, params.IDAlbum)
	if err != nil {
		return nil, err
	}
	result := make([]oas.Track, len(list.Items))
	for i, t := range list.Items {
		result[i] = oas.Track{ID: t.Id, Name: t.Name, IDAlbum: t.IdAlbum}
	}
	return result, nil
}

func (h *Handler) GetAuthorsFromAlbum(ctx context.Context, params oas.GetAuthorsFromAlbumParams) ([]oas.Author, error) {
	list, err := h.privateAppService.GetAuthorsByAlbum(ctx, params.IDAlbum)
	if err != nil {
		return nil, err
	}
	result := make([]oas.Author, len(list.Items))
	for i, a := range list.Items {
		result[i] = oas.Author{ID: a.Id, Username: a.Username, Name: a.Name}
	}
	return result, nil
}

func (h *Handler) AddAuthorToAlbum(ctx context.Context, req *oas.AddAuthorToAlbumReq, params oas.AddAuthorToAlbumParams) (*oas.AddAuthorToAlbumOK, error) {
	if err := h.privateAppService.UpdateAuthorAlbum(ctx, req.IDAuthor, params.IDAlbum); err != nil {
		return nil, err
	}
	return &oas.AddAuthorToAlbumOK{Success: oas.NewOptBool(true)}, nil
}

func (h *Handler) DeleteAuthorFromAlbum(ctx context.Context, params oas.DeleteAuthorFromAlbumParams) error {
	if err := h.privateAppService.RemoveAuthorAlbum(ctx, params.IDAuthor, params.IDAuthor); err != nil {
		return err
	}
	return nil
}
