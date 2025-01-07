package repository

import (
	"social-media-feed/internal/feed_data"

	"github.com/jmoiron/sqlx"
)

type PostGateway interface {
	Create(post feed_data.Post) (int, error)
	GetAll() ([]feed_data.Post, error)
	GetById(postId int) (feed_data.Post, error)
	Delete(postId int) error
	Update(postId int, input feed_data.PostUpdateFields) error
}

type LikeGateway interface {
	Create(like feed_data.Like) (int, error)
	GetAll(userId int) ([]feed_data.Like, error)
	GetById(userId, postId int) (feed_data.Like, error)
	Delete(userId, postId int) error
}

type Repository struct {
	PostGateway
	LikeGateway
	PostWithLikeGateway
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PostGateway:      NewPostGatewayPostgres(db),
		LikeGateway:      NewLikeGatewayPostgres(db),
		PostWithLikeGateway:	  NewPostLikeGatewayPostgres(db),
	}
}