package main

import (
	"net/http"
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
