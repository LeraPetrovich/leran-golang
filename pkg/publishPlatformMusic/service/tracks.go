package service

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (s *Service) GetAllTracks(ctx context.Context) (*core.TrackList, error) {
	return s.privateAppStorage.GetAllTracks(ctx)
}

func (s *Service) GetTrackById(ctx context.Context, trackId int) (*core.Track, error) {
	return s.privateAppStorage.GetTrackById(ctx, trackId)
}

func (s *Service) CreateNewTrack(ctx context.Context, track *core.Track, authorId int) (*core.Track, error) {
	return s.privateAppStorage.CreateNewTrack(ctx, track, authorId)
}

func (s *Service) UpdateTrack(ctx context.Context, idTrack int, track *core.Track) (*core.Track, error) {
	return s.privateAppStorage.UpdateTrack(ctx, idTrack, track)
}

func (s *Service) RemoveTrack(ctx context.Context, idTrack int) error {
	return s.privateAppStorage.RemoveTrack(ctx, idTrack)
}

func (s *Service) GetAllTracksByAuthor(ctx context.Context, authorId int) (*core.TrackList, error) {
	return s.privateAppStorage.GetAllTracksByAuthor(ctx, authorId)
}

func (s *Service) GetAlbumByTrack(ctx context.Context, trackId int) (*core.Album, error) {
	return s.privateAppStorage.GetAlbumByTrack(ctx, trackId)
}

func (s *Service) GetAuthorsByTrack(ctx context.Context, trackId int) (*core.AuthorsList, error) {
	return s.privateAppStorage.GetAuthorsByTrack(ctx, trackId)
}
