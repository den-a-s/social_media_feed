package repository

import (
	"errors"
	"social-media-feed/internal/feed_data"
    "fmt"
    "database/sql"
	"github.com/jmoiron/sqlx"
)

type LikeGatewayPostgres struct {
	db *sqlx.DB
}

func NewLikeGatewayPostgres(db *sqlx.DB) *LikeGatewayPostgres {
	return &LikeGatewayPostgres{db: db}
}

func (r *LikeGatewayPostgres) Create(like feed_data.Like) (int, error) {
	query := `INSERT INTO like (post_id, user_id) VALUES ($1, $2)`
	result, err := r.db.Exec(query, like.PostID, like.UserID)
    if err != nil {
        // log.Printf("Ошибка при вставке лайка: %v", err)
        return 0, errors.Wrap(err, "не удалось создать лайк")
    }
	insertedID, err := result.LastInsertId()
    if err != nil {
        // log.Printf("Ошибка при получении ID новой записи: %v", err)
        return 0, errors.Wrap(err, "не удалось получить ID новой записи")
    }
    return int(insertedID), nil
}

func (r *LikeGatewayPostgres) GetAll(userId int) ([]feed_data.Like, error) {
	likes := make([]feed_data.Like, 0)
	query := fmt.Sprintf(`SELECT id, post_id, user_id FROM like WHERE user_id = %d`, userId)
	if err := r.db.Select(&likes, query); err != nil {
		// log.Printf("Ошибка при получении лайков: %v", err)
        return nil, errors.Wrap(err, "не удалось получить лайки")
    }
    return likes, nil
}
        

func (r *LikeGatewayPostgres) GetById(userId, postId int) (feed_data.Like, error) {
    var like feed_data.Like
    query := `SELECT id, post_id, user_id FROM like WHERE user_id = $1 AND post_id = $2`
    if err := r.db.Get(&like, query, userId, postId); err != nil {
        if err == sql.ErrNoRows {
        	// log.Println("Лайк не найден")
             return feed_data.Like{}, sql.ErrNoRows
        }
        // log.Printf("Ошибка при получении лайка: %v", err)
        return feed_data.Like{}, err
    }
    return like, nil
}


func (r *LikeGatewayPostgres) Delete(userId, postId int) error {
	query := fmt.Sprintf(`DELETE FROM like WHERE user_id = %d AND post_id = %d`, userId, postId)
    result, err := r.db.Exec(query)
    if err != nil {
        return errors.Wrap(err, "не удалось удалить лайк")
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return errors.Wrap(err, "не удалось проверить количество удаленных строк")
    }
    if rowsAffected == 0 {
        // log.Println("Лайк не был найден для удаления")
        return sql.ErrNoRows
}
    return nil
}
