package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.indexHandler)
	mux.HandleFunc("GET /posts/create", app.createPost)
	mux.HandleFunc("POST /posts/create", app.storePost)
	mux.Handle("/styles/", http.StripPrefix("/styles", RestrictedFileServer(http.Dir("assets/styles"), []string{".css"})))
	return mux
}