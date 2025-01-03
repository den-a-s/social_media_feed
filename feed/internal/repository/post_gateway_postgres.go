package repository

import (
	"errors"
	"fmt"
	"social-media-feed/internal/feed_data"

	"github.com/jmoiron/sqlx"
)

type PostGatewayPostgres struct {
	db *sqlx.DB
}

func NewPostGatewayPostgres(db *sqlx.DB) *PostGatewayPostgres {
	return &PostGatewayPostgres{db: db}
}

func (r *PostGatewayPostgres) Create(post feed_data.Post) (int, error) {
	return 0, errors.New("TODO Implement method")
}

func (r *PostGatewayPostgres) GetAll() ([]feed_data.Post, error) {
	var posts []feed_data.Post
	query := fmt.Sprintf(`SELECT id, name, image_path, content FROM post`)
	if err := r.db.Select(&posts, query); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostGatewayPostgres) GetById(postId int) (feed_data.Post, error) {
	return feed_data.Post{}, errors.New("TODO Implement method")
}

func (r *PostGatewayPostgres) Delete(postId int) error {
	return errors.New("TODO Implement method")
}

func (r *PostGatewayPostgres) Update(postId int, input feed_data.PostUpdateFields) error {
	return errors.New("TODO Implement method")
}