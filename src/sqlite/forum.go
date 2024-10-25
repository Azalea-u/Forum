package sqlite

import (
	"database/sql"

	"forum/src/models"
)

type PostModel struct {
	DB *sql.DB
}

func (f *PostModel) Insert(title, content, category_id string) error{
	post := `INSERT INTO post (user_id, title, content, created_at, category_id)
	VALUES (1, ?, ?, datetime('now'), ?)`
	_, err := f.DB.Exec(post, title, content, category_id)
	return err
}

func (f *PostModel) Posts() ([]models.Post, error) {
	stmt := `SELECT p.id, p.user_id, p.title, p.content, p.created_at, p.category_id
		FROM Post p
		ORDER BY p.created_at DESC;`
	rows, err := f.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []models.Post{}
	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.CategoryID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
