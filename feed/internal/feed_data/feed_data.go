package feed_data

import "database/sql"

type Post struct {
	Id int `json:"id" db:"id" binding:"required"`
    ImagePath string `json:"image_path" db:"image_path" binding:"required"`
    Content sql.NullString `json:"content" db:"content"`
}

type PostUpdateFields struct {
    ImagePath string
    Content string
}

type Like struct {
	Id int `json:"id" db:"id" binding:"required"`
	PostId int `json:"post_id" db:"post_id" binding:"required"`
    UserId int `json:"user_id" db:"user_id" binding:"required"`
}

type PostWithLike struct {
    Id int `json:"id" db:"id" binding:"required"`
    ImagePath string `json:"image_path" db:"image_path" binding:"required"`
    Content sql.NullString `json:"content" db:"content"`
    LikeId sql.NullInt64 `json:"like_id" db:"like_id" binding:"required"`
    UserId sql.NullInt64 `json:"user_id" db:"user_id" binding:"required"`  
}