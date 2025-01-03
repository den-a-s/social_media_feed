package feed_data

import "database/sql"

type Post struct {
	Id int `json:"id" db:"id" binding:"required"`
    Name string `json:"name" db:"name" binding:"required"`
    ImagePath string `json:"image_path" db:"image_path"`
    Content sql.NullString `json:"content" db:"content"`
}

type PostUpdateFields struct {
    Name string
    ImagePath string
    Content string
}

type Like struct {
	Id int `json:"id" db:"id" binding:"required"`
	PostId int `json:"post_id" db:"post_id" binding:"required"`
    UserId int `json:"user_id" db:"user_id" binding:"required"`
}