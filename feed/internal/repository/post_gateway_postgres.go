package repository

import (
	"database/sql"
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
	var newID int
	query := `INSERT INTO post (image_path, content) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, post.ImagePath, post.Content).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func (r *PostGatewayPostgres) GetAll() ([]feed_data.Post, error) {
	var posts []feed_data.Post
	query := fmt.Sprintf(`SELECT id, image_path, content FROM post`)
	if err := r.db.Select(&posts, query); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostGatewayPostgres) GetById(postId int) (feed_data.Post, error) {
	query:=`SELECT * FROM post WHERE id = $1`
	
	var post feed_data.Post
    err := r.db.QueryRow(query, postId).Scan(&post.Id, &post.ImagePath, &post.Content) 

    if err != nil {
        if err == sql.ErrNoRows {
            return feed_data.Post{}, errors.New("post not found")
        }
        return feed_data.Post{}, err 
    }

    return post, nil 
}

func (r *PostGatewayPostgres) Delete(postId int) error {
	
    query := `DELETE FROM post WHERE id = $1`
    _, err := r.db.Exec(query, postId)
	return err
}
