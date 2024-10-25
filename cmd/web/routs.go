package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.indexHandler)
	mux.Handle("/styles/", http.StripPrefix("/styles", RestrictedFileServer(http.Dir("assets/styles"), []string{".css"})))
	return mux
}