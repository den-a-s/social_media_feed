package repository

import (
	"social-media-feed/internal/feed_data"
	"fmt"
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
	li.id AS like_id, 
	li.user_id AS user_id 
	FROM post p 
	RIGHT JOIN "like" li ON p.id = li.post_id AND li.user_id = $1;`)
	if err := r.db.Select(&postsWithLike, query, userId); err != nil {
		return nil, err
	}
	return postsWithLike, nil
}