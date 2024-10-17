package main

import (
	"database/sql"
	"fmt"
	forum "forum/src"
	"forum/src/sqlite"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	posts *sqlite.PostModel
}

func main() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Printf("%s Unable to open database. Reason: %v\n", forum.ERROR, err)
		return
	}
	app := app{
		posts: &sqlite.PostModel{DB: db},
	}
	defer db.Close()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", findValidPort()),
		Handler: app.routes(),
	}
	fmt.Printf("%s Starting server on: \033[1;32mhttp://localhost%s\033[0m\n", forum.INFO, srv.Addr)
	fmt.Println(forum.INFO, "Preparing to open the browser...")

	openBrowser(fmt.Sprintf("http://localhost%s", srv.Addr))

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("%s Unable to start server. Reason: %v\n", forum.ERROR, err)
	}
}