package repository

import (
	"errors"
	"social-media-feed/internal/feed_data"

	"github.com/jmoiron/sqlx"
)

type LikeGatewayPostgres struct {
	db *sqlx.DB
}

func NewLikeGatewayPostgres(db *sqlx.DB) *LikeGatewayPostgres {
	return &LikeGatewayPostgres{db: db}
}

func (r *LikeGatewayPostgres) Create(like feed_data.Like) (int, error) {
	return 0, errors.New("TODO Implement method")
}

func (r *LikeGatewayPostgres) GetAll(userId int) ([]feed_data.Like, error) {
	posts := make([]feed_data.Like, 0)
	return posts, errors.New("TODO Implement method")
}

func (r *LikeGatewayPostgres) GetById(userId, postId int) (feed_data.Like, error) {
	return feed_data.Like{}, errors.New("TODO Implement method")
}

func (r *LikeGatewayPostgres) Delete(userId, postId int) error {
	return errors.New("TODO Implement method")
}