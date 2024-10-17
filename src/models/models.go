package models

import "time"

// User represents the User table.
type User struct {
    ID          int
    Email       string
    Username    string
    PasswordHash string
    CreatedAt   time.Time
}

// Category represents the Category table.
type Category struct {
    ID   int
    Name string
}

// Post represents the Post table.
type Post struct {
    ID         int
    UserID     int
    Title      string
    Content    string
    CreatedAt  time.Time
    CategoryID *int
}

// Comment represents the Comment table.
type Comment struct {
    ID        int
    PostID    int
    UserID    int
    Content   string
    CreatedAt time.Time
}

// LikeDislike represents the LikeDislike table.
type LikeDislike struct {
    ID        int
    UserID    int
    PostID    *int
    CommentID *int
    Type      int
}
