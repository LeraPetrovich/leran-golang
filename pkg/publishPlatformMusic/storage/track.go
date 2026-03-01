package storage

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (s *storage) GetAllTracks(ctx context.Context) (*core.TrackList, error) {
	contextTime, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var tracks []*core.Track

	rows, err := s.postgres.Query(contextTime, `
	SELECT id, fk_id_album, name
	FROM tracks
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		track := &core.Track{}
		err := rows.Scan(&track.Id, &track.IdAlbum, &track.Name)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	return &core.TrackList{Items: tracks}, nil
}

func (s *storage) GetTrackById(ctx context.Context, trackId int) (*core.Track, error) {
	contextTime, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var track core.Track
	err := s.postgres.QueryRow(contextTime, `
    SELECT id, fk_id_album, name 
	FROM tracks
    WHERE id = $1
   `, trackId).Scan(&track.Id, &track.IdAlbum, &track.Name)

	if err != nil {
		return nil, err
	}
	return &track, nil
}

func (s *storage) GetTracksByAlbum(ctx context.Context, albumId int) (*core.TrackList, error) {
	contextTime, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var tracks []*core.Track

	rows, err := s.postgres.Query(contextTime, `
	SELECT id, fk_id_album, name
	FROM tracks
	WHERE fk_id_album = $1
	`, albumId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		track := &core.Track{}
		err := rows.Scan(&track.Id, &track.IdAlbum, &track.Name)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return &core.TrackList{Items: tracks}, nil
}

func (s *storage) GetAllTracksByAuthor(ctx context.Context, authorId int) (*core.TrackList, error) {
	contextTime, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var tracks []*core.Track

	rows, err := s.postgres.Query(contextTime, `
	SELECT t.id, t.fk_id_album, t.name
	FROM tracks t
	JOIN author_album aa ON t.fk_id_album = aa.album_id
	WHERE aa.author_id = $1
	`, authorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		track := &core.Track{}
		err := rows.Scan(&track.Id, &track.IdAlbum, &track.Name)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	return &core.TrackList{Items: tracks}, nil
}

func (s *storage) CreateNewTrack(ctx context.Context, track *core.Track, authorId int) (*core.Track, error) {
	contextTime, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var id int
	err := s.postgres.QueryRow(contextTime, `
	INSERT INTO tracks (name, fk_id_album)
	VALUES ($1, $2)
	RETURNING id
	`, track.Name, track.IdAlbum).Scan(&id)

	if err != nil {
		return nil, err
	}
	track.Id = id
	return track, nil
}

func (s *storage) UpdateTrack(ctx context.Context, idTrack int, track *core.Track) (*core.Track, error) {
	contextTime, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	//update track table
	cmdTag, err := s.postgres.Exec(contextTime, `
	UPDATE tracks
	SET name = $1, fk_id_album = $2
	WHERE id = $3
	`, track.Name, track.IdAlbum, idTrack)

	if err != nil {
		return nil, err
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, nil
	}
	track.Id = idTrack

	return track, nil
}

func (s *storage) RemoveTrack(ctx context.Context, idTrack int) error {
	contextTime, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	//delete track from track table
	_, err := s.postgres.Exec(contextTime, `
	DELETE FROM tracks
	WHERE id = $1
	`, idTrack)
	if err != nil {
		return err
	}
	return nil
}
