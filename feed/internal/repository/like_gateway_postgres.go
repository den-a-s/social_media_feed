package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"social-media-feed/internal/feed_data"

	"github.com/jmoiron/sqlx"
)

type LikeGatewayPostgres struct {
	db *sqlx.DB
}

func NewLikeGatewayPostgres(db *sqlx.DB) *LikeGatewayPostgres {
	return &LikeGatewayPostgres{db: db}
}

func (r *LikeGatewayPostgres) Create(like feed_data.Like) error {
	query := `INSERT INTO "like" (post_id, user_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, like.PostId, like.UserId)
	return err
}

// func (r *LikeGatewayPostgres) GetAll(userId int) ([]feed_data.Like, error) {
// 	likes := make([]feed_data.Like, 0)
// 	query := fmt.Sprintf(`SELECT id, post_id, user_id FROM like WHERE user_id = %d`, userId)
// 	if err := r.db.Select(&likes, query); err != nil {
// 		// log.Printf("Ошибка при получении лайков: %v", err)
//         return nil, errors.New(fmt.Sprintf("Error with getting likes: %s", err))
//     }
//     return likes, nil
// }

func (r *LikeGatewayPostgres) GetById(likeId int) ([]feed_data.Like, error) {
	var optional_like []feed_data.Like
	var like feed_data.Like
	query := `SELECT * FROM "like: WHERE id = $1`
	if err := r.db.Get(&like, query, likeId); err != nil {
		if err == sql.ErrNoRows {
			return optional_like, nil
		}
		return optional_like, err
	}
	optional_like = append(optional_like, like)
	return optional_like, nil
}

func (r *LikeGatewayPostgres) Delete(likeId int) error {
	query := fmt.Sprintf(`DELETE FROM "like" WHERE id = %d`, likeId)
	result, err := r.db.Exec(query)
	if err != nil {
		return errors.New(fmt.Sprintf("Error with deleting a like: %s", err))
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New(fmt.Sprintf("Error with checking number of deleted rows: %s", err))
	}
	if rowsAffected == 0 {
		// log.Println("Лайк не был найден для удаления")
		return sql.ErrNoRows
	}
	return nil
}
