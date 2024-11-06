package main

import (
	"net/http"
)

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.indexHandler)
	mux.HandleFunc("/post/", app.postDiscussion)

	mux.HandleFunc("/posts/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			app.createPost(w, r)
		} else if r.Method == http.MethodPost {
			app.storePost(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.Handle("/styles/", http.StripPrefix("/styles", RestrictedFileServer(http.Dir("assets/styles"), []string{".css"})))
	return mux
}
