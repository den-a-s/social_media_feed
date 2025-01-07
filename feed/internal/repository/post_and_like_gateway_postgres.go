package repository

import (
	"errors"
	"social-media-feed/internal/feed_data"

	"github.com/jmoiron/sqlx"
)

type PostLikeGatewayPostgres struct {
	db *sqlx.DB
}

func NewPostLikeGatewayPostgres(db *sqlx.DB) *PostLikeGatewayPostgres {
	return &PostLikeGatewayPostgres{db: db}
}


func (r *PostLikeGatewayPostgres) JoinPostWithLike(userId int) ([]feed_data.PostWithLike, error) {
	var postsWithLike []feed_data.PostWithLike
	query := fmt.Sprintf(`SELECT p.id AS id, 
	p.name AS name, 
	p.image_path AS image_path, 
	p.content AS content, 
	l.id AS like_id, 
	l.user_id AS user_id 
	FROM posts p 
	RIGHT JOIN likes l ON p.id = l.post_id AND l.user_id = $1;`)
	if err := r.db.Get(&postsWithLike, query, userId); err != nil {
		return nil, err
	}
	return postsWithLike, nil
}