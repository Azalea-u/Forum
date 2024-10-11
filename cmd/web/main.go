package main

import (
	"fmt"
	forum "forum/src"
	"net/http"
)

func main() {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", findValidPort()),
		Handler: routes(),
	}
	fmt.Printf("%s Starting server on: \033[1;32mhttp://localhost%s\033[0m\n", forum.INFO, srv.Addr)
	fmt.Println(forum.INFO, "Preparing to open the browser...")

	openBrowser(fmt.Sprintf("http://localhost%s", srv.Addr))

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("%s Unable to start server. Reason: %v\n", forum.ERROR, err)
	}
}
