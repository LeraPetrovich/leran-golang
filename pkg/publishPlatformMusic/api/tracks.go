package api

import (
	"context"
	"fmt"

	"example.com/learn/pkg/publishPlatformMusic/api/oas"
	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (h *Handler) GetAllTracks(ctx context.Context) ([]oas.Track, error) {
	list, err := h.privateAppService.GetAllTracks(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]oas.Track, len(list.Items))
	for i, t := range list.Items {
		result[i] = oas.Track{ID: t.Id, Name: t.Name, IDAlbum: t.IdAlbum}
	}
	return result, nil
}

func (h *Handler) GetTrackById(ctx context.Context, params oas.GetTrackByIdParams) (*oas.Track, error) {
	t, err := h.privateAppService.GetTrackById(ctx, params.IDTrack)
	if err != nil {
		return nil, err
	}
	return &oas.Track{ID: t.Id, Name: t.Name, IDAlbum: t.IdAlbum}, nil
}

func (h *Handler) CreateNewTrack(ctx context.Context, req *oas.CreateNewTrackReq) (*oas.Track, error) {
	claims, ok := ctx.Value(UserContextKey).(*Claims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	track := &core.Track{Name: req.Name, IdAlbum: req.IDAlbum}
	created, err := h.privateAppService.CreateNewTrack(ctx, track, claims.UserId)
	if err != nil {
		return nil, err
	}
	return &oas.Track{ID: created.Id, Name: created.Name, IDAlbum: created.IdAlbum}, nil
}

func (h *Handler) UpdateTrack(ctx context.Context, req *oas.UpdateTrackReq, params oas.UpdateTrackParams) (*oas.Track, error) {
	track := &core.Track{Name: req.Name, IdAlbum: req.IDAlbum}
	updated, err := h.privateAppService.UpdateTrack(ctx, params.IDTrack, track)
	if err != nil {
		return nil, err
	}
	if updated == nil {
		return nil, fmt.Errorf("track not found")
	}
	return &oas.Track{ID: updated.Id, Name: updated.Name, IDAlbum: updated.IdAlbum}, nil
}

func (h *Handler) DeleteTrack(ctx context.Context, params oas.DeleteTrackParams) error {
	return h.privateAppService.RemoveTrack(ctx, params.IDTrack)
}

func (h *Handler) GetAlbumFromTrack(ctx context.Context, params oas.GetAlbumFromTrackParams) (*oas.Album, error) {
	a, err := h.privateAppService.GetAlbumByTrack(ctx, params.IDTrack)
	if err != nil {
		return nil, err
	}
	return &oas.Album{ID: a.Id, Title: a.Title, Description: a.Description}, nil
}

func (h *Handler) GetAuthorsFromTrack(ctx context.Context, params oas.GetAuthorsFromTrackParams) ([]oas.Author, error) {
	list, err := h.privateAppService.GetAuthorsByTrack(ctx, params.IDTrack)
	if err != nil {
		return nil, err
	}
	result := make([]oas.Author, len(list.Items))
	for i, a := range list.Items {
		result[i] = oas.Author{ID: a.Id, Username: a.Username, Name: a.Name}
	}
	return result, nil
}
