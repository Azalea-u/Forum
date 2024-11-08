package sqlite

import (
	"database/sql"
	"forum/src/models"
)

type PostModel struct {
	DB *sql.DB
}

// Insert a new post into the Post table
func (f *PostModel) Insert(title, content, categoryID string) error {
	query := `INSERT INTO Post (user_id, title, content, created_at, category_id)
	          VALUES (1, ?, ?, datetime('now'), ?)`
	_, err := f.DB.Exec(query, title, content, categoryID)
	return err
}

// Posts retrieves all posts in descending order of creation
func (f *PostModel) Posts() ([]models.Post, error) {
	stmt := `SELECT p.id, p.user_id, p.title, p.content, p.created_at, p.category_id
	         FROM Post p
	         ORDER BY p.created_at DESC`
	rows, err := f.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		// Use pointers to handle nullable fields, e.g., CategoryID
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.CategoryID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (f *PostModel) GetByID(id int) (models.Post, []models.Comment, error) {
	var post models.Post
	query := `SELECT id, user_id, title, content, created_at, category_id FROM Post WHERE id = ?`

	// Retrieve post
	err := f.DB.QueryRow(query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.CategoryID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return post, nil, nil // No post found
		}
		return post, nil, err
	}

	// Retrieve comments
	commentQuery := `SELECT id, post_id, user_id, content, created_at FROM Comment WHERE post_id = ?`
	rows, err := f.DB.Query(commentQuery, id)
	if err != nil {
		return post, nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt); err != nil {
			return post, nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return post, nil, err
	}

	return post, comments, nil
}
