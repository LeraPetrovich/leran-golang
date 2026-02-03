package storage

import (
	"context"

	"example.com/learn/pkg/publishPlatformMusic/core"
)

func (s *storage) GetAllAlbums(ctx context.Context) (*core.AlbumsList, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var albums []*core.Album

	rows, err := s.postgres.Query(contextWait, `
	SELECT id, title, description
	FROM albums
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album core.Album
		err := rows.Scan(&album.Id, &album.Title, &album.Description)
		if err != nil {
			return nil, err
		}
		albums = append(albums, &album)
	}

	return &core.AlbumsList{Items: albums}, nil
}

func (s *storage) GetAlbumFromId(ctx context.Context, albumId int) (*core.Album, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var album core.Album
	err := s.postgres.QueryRow(contextWait, `
	SELECT id, title, description
	FROM albums
	WHERE id = $1
	`, albumId).Scan(&album.Id, &album.Title, &album.Description)

	if err != nil {
		return nil, err
	}

	return &album, err
}

func (s *storage) GetAlbumFromTrackId(ctx context.Context, trackId int) (*core.Album, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var album core.Album

	err := s.postgres.QueryRow(contextWait, `
	SELECT a.id, a.title, a.description
	FROM albums a
	JOIN tracks t ON t.fk_id_album = a.id
	WHERE t.id = $1
	`, trackId).Scan(&album.Id, &album.Title, &album.Description)

	if err != nil {
		return nil, err
	}

	return &album, nil
}

func (s *storage) GetAllAlbumsFromAuthor(ctx context.Context, authorId int) (*core.AlbumsList, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var albums []*core.Album

	rows, err := s.postgres.Query(contextWait, `
	SELECT a.id, a.title, a.description
	FROM albums a
	JOIN author_album aa ON a.id = aa.album_id
	WHERE aa.author_id = $1
	`, authorId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album core.Album
		err := rows.Scan(&album.Id, &album.Title, &album.Description)
		if err != nil {
			return nil, err
		}
		albums = append(albums, &album)
	}

	return &core.AlbumsList{Items: albums}, nil
}

func (s *storage) CreateNewAlbum(ctx context.Context, album *core.Album, authorId int) (*core.Album, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	var id int
	err := s.postgres.QueryRow(contextWait, `
	INSERT INTO albums (title, description)
	VALUES ($1, $2)
	RETURNING id
	`, album.Title, album.Description).Scan(&id)

	if err != nil {
		return nil, err
	}

	errAlbumAuthor := s.postgres.QueryRow(contextWait, `
	INSERT INTO author_album (album_id, author_id)
	VALUES ($1, $2)
	RETURNING album_id
	`, id, authorId).Scan(&id)

	if errAlbumAuthor != nil {
		return nil, errAlbumAuthor
	}

	album.Id = id

	return album, nil
}

func (s *storage) UpdateAlbum(ctx context.Context, idAlbum int, album *core.Album) (*core.Album, error) {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	cmdTag, err := s.postgres.Exec(contextWait, `
	UPDATE albums
	SET title = $1, description = $2
	WHERE id = $3
	`, album.Title, album.Description, idAlbum)

	if err != nil {
		return nil, err
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, nil
	}

	album.Id = idAlbum

	return album, nil
}

func (s *storage) RemoveAlbum(ctx context.Context, idAlbum int) error {
	contextWait, cancel := context.WithTimeout(ctx, fastQueryTimeout)
	defer cancel()

	_, errAlbumAuthor := s.postgres.Exec(contextWait, `
	DELETE FROM author_album
	WHERE album_id = $1
	`, idAlbum)

	if errAlbumAuthor != nil {
		return errAlbumAuthor
	}

	_, errTrack := s.postgres.Exec(contextWait, `
	UPDATE tracks
	SET fk_id_album = NULL
	WHERE fk_id_album = $1
	`, idAlbum)

	if errTrack != nil {
		return errTrack
	}

	_, errAlbums := s.postgres.Exec(contextWait, `
	DELETE FROM albums
	WHERE id = $1
	`, idAlbum)

	return errAlbums

}
