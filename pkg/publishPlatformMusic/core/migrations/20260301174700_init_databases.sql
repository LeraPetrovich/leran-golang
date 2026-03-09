-- +goose Up
CREATE TABLE authors(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username VARCHAR(200) NOT NULL,
    name VARCHAR(200) NOT NULL,
    refresh_token TEXT
);

INSERT INTO authors (username, name, refresh_token)
VALUES ('lera_test', 'Valeria', null);

CREATE TABLE albums(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description VARCHAR(500)
);

INSERT INTO albums (title, description)
VALUES ('lera_album', 'First Valeria Release');

CREATE TABLE tracks(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    fk_id_album int REFERENCES albums(id)
);

CREATE TABLE author_album(
    album_id int NOT NULL,
    author_id int NOT NULL,
    CONSTRAINT unique_album_author UNIQUE (album_id, author_id),
    FOREIGN KEY (album_id) REFERENCES albums(id),
    FOREIGN KEY (author_id) REFERENCES authors(id)
);
-- +goose Down
DROP TABLE author_album;
DROP TABLE tracks;
DROP TABLE authors;
DROP TABLE albums;
