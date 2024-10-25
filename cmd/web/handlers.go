package main

import (
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
)

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

// RestrictedFileServer serves only files with the specified extensions.
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
