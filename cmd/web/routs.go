package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.indexHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets", RestrictedFileServer(http.Dir("assets"), []string{".css", ".html",".js"})))
	return mux
}