package main

import (
	"forum/src/models"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

// Handler for the index page, listing posts
func (app *app) indexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := app.posts.Posts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("./assets/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, map[string]any{"Posts": posts}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handler for rendering the create post form
func (app *app) createPost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/post.create.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handler for processing the form submission to store a post
func (app *app) storePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = app.posts.Insert(
		r.PostForm.Get("title"),
		r.PostForm.Get("content"),
		r.PostForm.Get("category_id"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// Handler for displaying a single post based on post ID
func (app *app) postDiscussion(w http.ResponseWriter, r *http.Request) {
	// Extract post ID from URL
	path := strings.TrimPrefix(r.URL.Path, "/post/")
	postID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// Retrieve the post by ID
	post, err := app.posts.GetByID(postID)
	if err != nil {
		http.Error(w, "Unable to retrieve post", http.StatusInternalServerError)
		return
	}
	if post.ID == 0 { // Checking if post was found
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	// Load and render the template with the retrieved post
	t, err := template.ParseFiles("./assets/templates/post.discussion.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Post models.Post
	}{
		Post: post,
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// RestrictedFileServer serves only files with the specified extensions
func RestrictedFileServer(fs http.FileSystem, allowedExtensions []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := strings.ToLower(filepath.Ext(r.URL.Path))
		for _, allowedExt := range allowedExtensions {
			if ext == allowedExt {
				http.FileServer(fs).ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "Access forbidden", http.StatusForbidden)
	})
}
